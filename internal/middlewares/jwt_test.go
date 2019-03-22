package middlewares

import (
	"fmt"
	"testing"

	"github.com/dgrijalva/jwt-go"
)

var jwtMid = JWTMiddleware{
	key:           []byte("aasdasfqw"),
	signingMethod: jwt.SigningMethodHS256,
}

func Test_JWTMiddleware_signToken(t *testing.T) {
	fmt.Println(jwtMid.signToken(&payloadJWT{}))
}
func Test_JWTMiddleware_verifyToken(t *testing.T) {
	fmt.Println(jwtMid.verifyToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoid3pzIn0.821WO5tagyM8dKQpZ8fms4seR19XNus_vwfN4TqrP0g"))
}
