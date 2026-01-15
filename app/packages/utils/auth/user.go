package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func getJWTSigningKey() []byte {
	key := os.Getenv("JWT_SIGNING_KEY")
	if key == "" {
		// 開発環境用のデフォルトキー（本番では必ず環境変数を設定）
		key = "SKNGIonriongINIOnfiar394rjOJGg"
	}
	return []byte(key)
}

var userTokenJwt = &hs256jwt{
	sigKey: getJWTSigningKey(),
	createClaims: func() jwt.Claims {
		return &userClaims{}
	},
}

type userClaims struct {
	Subject  string `json:"user_id"`
	IssuedAt int64  `json:"iat"`
	Exp      int64  `json:"exp"`
}

var mobileUserTokenJwt = &hs256jwtMobile{
	sigKey: getJWTSigningKey(),
	createClaims: func() jwt.Claims {
		return &mobileUserClaims{}
	},
}

type mobileUserClaims struct {
	Subject  string `json:"user_id"`
	IssuedAt int64  `json:"iat"`
}

func (c *userClaims) Valid() error {
	now := time.Now()
	if c.IssuedAt > now.Unix() {
		return fmt.Errorf("issued on future time: %d (now:%d)", c.IssuedAt, now.Unix())
	}

	return nil
}

func IssueUserToken(userID string) (string, error) {

	now := time.Now()

	claims := &userClaims{
		Subject:  userID,
		IssuedAt: now.Unix(),
		Exp:      now.Add(30 * time.Minute).Unix(),
	}

	return userTokenJwt.issueToken(claims)
}

func VerifyUserToken(tokenStr string) (string, error) {
	claims, err := userTokenJwt.verifyToken(tokenStr)
	if err != nil {
		return "", err
	}
	return claims.(*userClaims).Subject, nil
}

func (c *mobileUserClaims) Valid() error {
	now := time.Now()
	if c.IssuedAt > now.Unix() {
		return fmt.Errorf("issued on future time: %d (now:%d)", c.IssuedAt, now.Unix())
	}

	return nil
}

func IssueMobileUserToken(userID string) (string, error) {

	now := time.Now()

	claims := &mobileUserClaims{
		Subject:  userID,
		IssuedAt: now.Unix(),
	}

	return mobileUserTokenJwt.issueMobileToken(claims)
}

func VerifyMobileUserToken(tokenStr string) (string, error) {
	claims, err := mobileUserTokenJwt.verifyMobileToken(tokenStr)
	if err != nil {
		return "", err
	}
	return claims.(*mobileUserClaims).Subject, nil
}
