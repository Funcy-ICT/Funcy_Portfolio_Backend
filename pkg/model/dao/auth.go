package dao

import (
	"backend/pkg/model/dto"
	"backend/pkg/utils"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
)

const (
	AuthenticationMail  = "SELECT COUNT(*) `users`  FROM `users` WHERE `mail`=? ;"
	InsertUserInfoQuery = "INSERT INTO `users` (id,display_name,icon,family_name,first_name,mail,password,grade,course,token) VALUES (?,?,?,?,?,?,?,?,?,?)"
	AuthenticatioToken  = "SELECT password FROM `users` WHERE `mail`=? ;"
	UpdateToken         = "UPDATE `users` set `token` = ? where `mail` = ? ;"
)

// sign/up
type signUp struct {
}

func MakeSignUpClient() signUp {
	return signUp{}
}

func (info *signUp) Request(userInfo dto.SignUpRequest) (string, error) {
	var exsit int

	userID, err := uuid.NewRandom()

	log.Println(userID)
	if err != nil {
		return "", errors.New("userID generate is failed")
	}

	password, err := utils.PasswordEncrypt(userInfo.Password)
	if err != nil {
		return "", errors.New("password generate is failed")
	}

	t, err := uuid.NewRandom()
	if err != nil {
		return "", errors.New("tokenID generate is failed")
	}
	token := t.String()

	row := Conn.QueryRow(AuthenticationMail, userInfo.Mail)
	if err = row.Scan(&exsit); err != nil {
		//if err == sql.ErrNoRows {
		//	return "", errors.New("Already registered with that email address")
		//}
		log.Println(err)
	}

	//メアドがすでに使用されているかチェック
	//TODO　存在する学内メールアドレスか検証する処理を追加
	if exsit != 0 {
		log.Println(exsit)
		return "", errors.New("Already registered with that email address")
	}

	stmt, err := Conn.Prepare(InsertUserInfoQuery)
	if err != nil {
		return "", err
	}
	_, err = stmt.Exec(userID, userInfo.DisplayName, userInfo.Icon, userInfo.FamilyName, userInfo.FirstName, userInfo.Mail, password, userInfo.Grade, userInfo.Course, token)

	return token, err
}

// sign/in
type signIn struct {
}

func MakeSignInClient() signIn {
	return signIn{}
}

func (info *signIn) Request(userInfo dto.SignInRequest) (string, error) {
	var hashPass string

	row := Conn.QueryRow(AuthenticatioToken, userInfo.Mail)
	if err := row.Scan(&hashPass); err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("mail address is not true")
		}
		log.Println(err)
	}
	// ハッシュ値でのパスワード比較
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(userInfo.Password))
	if err != nil {
		log.Println(err)
		return "", errors.New("not match password")
	}

	token := uuid.NewString()
	if err != nil {
		log.Println("tokenID is refresh")
	}

	stmt, err := Conn.Prepare(UpdateToken)
	if err != nil {
		return "", errors.New("not update")
	}

	_, err = stmt.Exec(token, userInfo.Mail)
	if err != nil {
		return "", errors.New("not update")
	}
	return token, err
}
