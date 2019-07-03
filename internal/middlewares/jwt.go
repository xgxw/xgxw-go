package middlewares

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	flog "github.com/everywan/foundation-go/log"
	"github.com/everywan/xgxw/internal/constants"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

/*
	后端是否有login路由有两种处理逻辑
	当使用cookie存储token时, 可以不需要login路由, 因为认证中间件中可以直接设置cookie.
	当使用localstorage/queryParam时, 需要login路由, 因为需要写入到response.data中.

	当前端有登录页时, 需要login路由, 否则登录页没有合适的请求地址. 并且前端可以不判断 未认证 错误, 因为后端可以直接重定向到登录页. 但是不建议这么做, 因为跳转前最好通知用户确认.
	当不使用登录页时, 可以在页面弹窗提示输入认证信息, 且必须验证后端返回值有无 未认证 错误.

	目前采用的方案
		1. 无login路由(即使用cookie存储token)
		2. 在前端每个路由请求处, 判断错误信息, 且判断未认证错误.
*/

const (
	// DefaultExpires 默认存活时间
	DefaultExpires = time.Hour * 2
)

// 如何复用 jwtmiddleware 这一部分? 应该是通用的
// 如何更好的存储 Expires 这个字段
type (
	// AuthenticationOptions 认证中间件配置项
	AuthenticationOptions struct {
		Key     string `json:"-" yaml:"key" mapstructure:"key"`
		Expires int64  `json:"-" yaml:"expires" mapstructure:"expires"`
		Cipher  string `json:"-" yaml:"cipher" mapstructure:"cipher"`
	}
	// JWTMiddleware : only use HMAC
	JWTMiddleware struct {
		key           []byte
		signingMethod jwt.SigningMethod
		parse         *jwt.Parser
		expires       time.Duration // 用于StandardClaims
		logger        *flog.Logger
		cipher        string
	}
)

type (
	payloadClaims struct {
		UserID uint `json:"user_id" mapstructure:"user_id"`
		*jwt.StandardClaims
	}
)

// NewJWTMiddlewares 生成JWT中间件
func NewJWTMiddlewares(logger *flog.Logger, opts AuthenticationOptions) echo.MiddlewareFunc {
	jwtPrase := new(jwt.Parser)
	key := []byte(opts.Key)
	expires := DefaultExpires
	if opts.Expires != 0 {
		expires = time.Duration(opts.Expires)
	}
	jwt := &JWTMiddleware{
		key:           key,
		signingMethod: jwt.GetSigningMethod("HS256"),
		parse:         jwtPrase,
		expires:       expires,
		logger:        logger,
		cipher:        opts.Cipher,
	}
	return jwt.middlewareFunc
}

func (a *JWTMiddleware) middlewareFunc(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return a.handle(c, next)
	}
}

func (a *JWTMiddleware) handle(c echo.Context, next echo.HandlerFunc) error {
	var payload *payloadClaims
	cookie, err := c.Cookie(constants.Token)
	// 如果cookie中没有token信息, 则判断request中是否有认证信息. 反之直接取token校验
	if err != nil || cookie.Value == "" {
		var ok bool
		if payload, ok = a.verifyAuth(c); !ok {
			return a.redirectAuth(c)
		}
		// 有认证信息, 则签名payload并写入cookie
		tokenStr, _ := a.signToken(payload)
		cookie := &http.Cookie{
			Name:     constants.Token,
			Value:    tokenStr,
			Expires:  time.Now().Add(a.expires),
			HttpOnly: true,
		}
		c.SetCookie(cookie)
	} else {
		tokenStr := cookie.Value
		payload, err = a.verifyToken(tokenStr)
		if err != nil {
			a.logger.Fatalf("verifyToken error: %v", err)
			return a.redirectAuth(c)
		}
	}

	c.Set(constants.UserID, payload.UserID)
	return next(c)
}

func (a *JWTMiddleware) verifyAuth(c echo.Context) (payload *payloadClaims, ok bool) {
	payload = new(payloadClaims)
	cipher := c.QueryParam(constants.Cipher)
	if cipher != a.cipher {
		return payload, false
	}
	payload = a.genPayloadClaims(1)
	return payload, true
}

func (a *JWTMiddleware) redirectAuth(c echo.Context) error {
	return c.NoContent(http.StatusForbidden)
}

func (a *JWTMiddleware) genPayloadClaims(userID uint) *payloadClaims {
	standClaims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(a.expires).Unix(),
	}
	payload := &payloadClaims{
		UserID:         userID,
		StandardClaims: standClaims,
	}
	return payload
}

func (a *JWTMiddleware) signToken(claims jwt.Claims) (tokenStr string, err error) {
	token := jwt.NewWithClaims(a.signingMethod, claims)
	out, err := token.SignedString(a.key)
	return out, err
}

func (a *JWTMiddleware) verifyToken(tokenStr string) (payload *payloadClaims, err error) {
	token, err := a.parse.ParseWithClaims(tokenStr, &payloadClaims{}, func(t *jwt.Token) (interface{}, error) {
		return a.key, nil
	})
	if err != nil {
		return nil, err
	}
	if payload, ok := token.Claims.(*payloadClaims); ok && token.Valid {
		return payload, nil
	}
	return nil, errors.Wrap(err, "parse jwt token error")
}
