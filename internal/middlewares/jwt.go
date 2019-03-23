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

const (
	DefaultExpires = time.Hour * 2
)

type (
	// AuthenticationOptions 认证中间件配置项
	AuthenticationOptions struct {
		Key     string `json:"-" yaml:"key" mapstructure:"-"`
		Expires int64  `json:"-" yaml:"expires" mapstructure:"-"`
	}
	// JWTMiddleware : only use HMAC
	JWTMiddleware struct {
		key           []byte
		signingMethod jwt.SigningMethod
		parse         *jwt.Parser
		expires       time.Duration
		logger        *flog.Logger
	}
	payloadClaims struct {
		UserID uint `json:"user_id",mapstructure:"user_id"`
		*jwt.StandardClaims
	}
)

var (
	standClaims = new(jwt.StandardClaims)
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
	}
	return jwt.MiddlewareFunc
}

func (a *JWTMiddleware) MiddlewareFunc(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return a.Handle(c, next)
	}
}

func (a *JWTMiddleware) Handle(c echo.Context, next echo.HandlerFunc) error {
	payload := new(payloadClaims)
	cookie, err := c.Cookie(constants.Token)
	if err != nil || cookie.Value == "" {
		// 判断是否是OA授权调回
		if true {
			expires := time.Now().Add(a.expires)
			standClaims.ExpiresAt = expires.Unix()
			// 假设成功, 且payload为获取的信息
			payload = &payloadClaims{
				UserID:         1,
				StandardClaims: standClaims,
			}
			tokenStr, err := a.signToken(payload)
			if err != nil {
				return err
			}
			cookie := &http.Cookie{
				Name:     constants.Token,
				Value:    tokenStr,
				Expires:  expires,
				HttpOnly: true,
			}
			c.SetCookie(cookie)
		} else {
			return a.redirectAuth(c)
		}
	} else {
		tokenStr := cookie.Value
		payload, err = a.verifyToken(tokenStr)
		if err != nil {
			a.logger.Fatalf("verifyToken error: %v", err)
			return a.redirectAuth(c)
		}
	}
	payload.setContext(c)
	return next(c)
}

func (a *JWTMiddleware) redirectAuth(c echo.Context) error {
	// 跳转到OA认证
	url := ""
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (a *JWTMiddleware) signToken(claims *payloadClaims) (tokenStr string, err error) {
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

func (p *payloadClaims) setContext(c echo.Context) {
	if p.UserID != 0 {
		c.Set(constants.UserID, p.UserID)
	}
}
