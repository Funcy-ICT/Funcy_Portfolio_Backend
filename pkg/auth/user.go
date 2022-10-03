package auth

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

var userTokenJwt = &hs256jwt{
	sigKey: []byte("SKNGIonriongINIOnfiar394rjOJGg"),
	createClaims: func() jwt.Claims {
		return &userClaims{}
	},
}

type userClaims struct {
	Subject  int64 `json:"user_id"`
	IssuedAt int64 `json:"iat"`
	//Exp      int64 `json:"exp"`
}

func (c *userClaims) Valid() error {
	now := time.Now()
	if c.IssuedAt > now.Unix() {
		return fmt.Errorf("issued on future time: %d (now:%d)", c.IssuedAt, now.Unix())
	}

	return nil
}

func IssueUserToken(userID int64) (string, error) {

	now := time.Now()

	claims := &userClaims{
		Subject:  userID,
		IssuedAt: now.Unix(),
	}

	return userTokenJwt.issueToken(claims)
}

func VerifyUserToken(tokenStr string) (int64, error) {
	claims, err := userTokenJwt.verifyToken(tokenStr)
	if err != nil {
		return 0, err
	}

	return claims.(*userClaims).Subject, nil
}
