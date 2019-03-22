package middlewares

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

type (
	// AuthenticationOptions 认证中间件配置项
	AuthenticationOptions struct {
		Key string `json:"-" yaml:"key" mapstructure:"-"`
	}
	// JWTMiddleware : only use HMAC
	JWTMiddleware struct {
		key           []byte
		signingMethod jwt.SigningMethod
	}
	payloadJWT struct{}
)

func (a *JWTMiddleware) MiddlewareFunc(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return a.Handle(c, next)
	}
}

func (a *JWTMiddleware) Handle(c echo.Context, next echo.HandlerFunc) error {
	payload := &payloadJWT{}
	var err error
	tokenStr := c.Request().Header.Get("Authorization")
	if tokenStr == "" {
		tokenStr, err = a.signToken(payload)
		if err != nil {
			return err
		}
		return a.redirectAuth(c, tokenStr)
	}
	payload, err = a.verifyToken(tokenStr)
	fmt.Println(payload)
	return err
}

func (a *JWTMiddleware) signToken(payload *payloadJWT) (tokenStr string, err error) {
	claims := jwt.MapClaims{}
	token := jwt.NewWithClaims(a.signingMethod, claims)
	out, err := token.SignedString(a.key)
	return out, err
}

func (a *JWTMiddleware) verifyToken(tokenStr string) (payload *payloadJWT, err error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return a.key, nil
	})
	if payload, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(payload)
		return &payloadJWT{}, nil
	}
	return nil, errors.Wrap(err, "parse jwt token error")
}

func (a *JWTMiddleware) redirectAuth(c echo.Context, tokenStr string) error {
	return nil
	// return c.Redirect(http.StatusTemporaryRedirect, "")
}

func NewJWTMiddlewares(opts AuthenticationOptions) (echo.MiddlewareFunc, error) {
	key := []byte(opts.Key)

	jwt := &JWTMiddleware{
		key:           key,
		signingMethod: &jwt.SigningMethodHMAC{},
	}
	return jwt.MiddlewareFunc, nil
}
