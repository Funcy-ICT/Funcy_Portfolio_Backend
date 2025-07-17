package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	RefreshTokenType = "refresh"
)

var userTokenJwt = &hs256jwt{
	sigKey: []byte("SKNGIonriongINIOnfiar394rjOJGg"),
	createClaims: func() jwt.Claims {
		return &userClaims{}
	},
}

type userClaims struct {
	Subject  string `json:"user_id"`
	IssuedAt int64  `json:"iat"`
	Exp      int64  `json:"exp"`
}

type refreshTokenClaims struct {
	Subject   string `json:"user_id"`
	IssuedAt  int64  `json:"iat"`
	Exp       int64  `json:"exp"`
	TokenType string `json:"token_type"`
}

var refreshTokenJwt = &hs256jwt{
	sigKey: []byte("SKNGIonriongINIOnfiar394rjOJGg"),
	createClaims: func() jwt.Claims {
		return &refreshTokenClaims{}
	},
}

var mobileUserTokenJwt = &hs256jwtMobile{
	sigKey: []byte("SKNGIonriongINIOnfiar394rjOJGg"),
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
	if c.Exp < now.Unix() {
		return fmt.Errorf("token expired: %d (now:%d)", c.Exp, now.Unix())
	}

	return nil
}

func (c *refreshTokenClaims) Valid() error {
	now := time.Now()
	if c.IssuedAt > now.Unix() {
		return fmt.Errorf("issued on future time: %d (now:%d)", c.IssuedAt, now.Unix())
	}
	if c.Exp < now.Unix() {
		return fmt.Errorf("token expired: %d (now:%d)", c.Exp, now.Unix())
	}
	if c.TokenType != RefreshTokenType {
		return fmt.Errorf("invalid token type: %s", c.TokenType)
	}

	return nil
}

func IssueUserToken(userID string) (string, error) {

	now := time.Now()

	claims := &userClaims{
		Subject:  userID,
		IssuedAt: now.Unix(),
		Exp:      now.Add(2 * time.Hour).Unix(), // 2時間に短縮（リフレッシュトークンと併用）
	}

	return userTokenJwt.issueToken(claims)
}

func IssueRefreshToken(userID string) (string, error) {
	now := time.Now()

	claims := &refreshTokenClaims{
		Subject:   userID,
		IssuedAt:  now.Unix(),
		Exp:       now.Add(7 * 24 * time.Hour).Unix(), // 7日間
		TokenType: RefreshTokenType,
	}

	return refreshTokenJwt.issueToken(claims)
}

func VerifyUserToken(tokenStr string) (string, error) {
	claims, err := userTokenJwt.verifyToken(tokenStr)
	if err != nil {
		return "", err
	}
	return claims.(*userClaims).Subject, nil
}

func VerifyRefreshToken(tokenStr string) (string, error) {
	claims, err := refreshTokenJwt.verifyToken(tokenStr)
	if err != nil {
		return "", err
	}
	return claims.(*refreshTokenClaims).Subject, nil
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
