package auth

import (
	"fmt"

	jwt "github.com/golang-jwt/jwt/v4"
)

type hs256jwt struct {
	sigKey       []byte
	createClaims func() jwt.Claims
}

var hs256jwtParser = &jwt.Parser{
	ValidMethods: []string{jwt.SigningMethodHS256.Alg()},
}

func (t *hs256jwt) issueToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(t.sigKey)
}

func (t *hs256jwt) verifyToken(tokenStr string) (jwt.Claims, error) {
	token, err := hs256jwtParser.ParseWithClaims(tokenStr, t.createClaims(), jwt.Keyfunc(func(token *jwt.Token) (interface{}, error) {
		return t.sigKey, nil
	}))
	if err != nil {
		return nil, fmt.Errorf("failed to verify JWT: %w", err)
	}

	return token.Claims, nil
}
