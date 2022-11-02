package entity

import (
	"backend/app/interfaces/request"
)

type Token struct {
	Token string `json:"token"`
}

type User struct {
	UserID      string `db:"id"`
	Token       string `db:"token"`
	Icon        string `db:"icon"`
	FamilyName  string `db:"family_name"`
	FirstName   string `db:"first_name"`
	Mail        string `db:"mail"`
	Password    string `db:"password"`
	Grade       string `db:"grade"`
	Course      string `db:"course"`
	DisplayName string `db:"display_name"`
}

func NewUser(user *request.SignUpRequest, userID, token string) (*User, error) {
	body := User{
		UserID:      userID,
		Token:       token,
		Icon:        user.Icon,
		FamilyName:  user.FamilyName,
		FirstName:   user.FirstName,
		Mail:        user.Mail,
		Password:    user.Password,
		Grade:       user.Grade,
		Course:      user.Course,
		DisplayName: user.DisplayName,
	}
	return &body, nil
}
