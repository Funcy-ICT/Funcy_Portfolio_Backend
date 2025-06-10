package usecase

import (
	"backend/app/domain/entity"
	"backend/app/domain/repository"
	"backend/app/interfaces/request"
	"backend/app/packages/utils"
	"backend/app/packages/utils/auth"
	"backend/app/packages/utils/mail"
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"

	"errors"
)

type AuthUseCase struct {
	authRepository repository.AuthRepository
}

func NewAuthUseCase(authRepository repository.AuthRepository) *AuthUseCase {
	return &AuthUseCase{authRepository: authRepository}
}

func (a *AuthUseCase) CreateAccount(r request.SignUpRequest) (string, error) {

	userID, err := uuid.NewRandom()
	if err != nil {
		return "", errors.New("userID generate is failed")
	}
	r.Password, err = utils.PasswordEncrypt(r.Password)
	if err != nil {
		return "", errors.New("password generate is failed")
	}
	token, err := uuid.NewRandom()
	if err != nil {
		return "", errors.New("tokenID generate is failed")
	}

	userStatus := "inactive"

	var authCode string
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		r := rand.Intn(9)
		authCode = authCode + strconv.Itoa(r)
	}

	user, err := entity.NewUser(&r, userID.String(), token.String(), authCode, userStatus)
	if err != nil {
		return "", err
	}

	err = a.authRepository.InsertAccount(user)
	if err != nil {
		return "", err
	}
	//mail本文生成
	text := "認証コード:" + authCode + "\n\n もしお心当たりのない場合、本メールは破棄して頂けるようお願いいたします \n\n *このメールへの返信はできません。"
	html := "認証コード:" + authCode + "<br><br>もしお心当たりのない場合、本メールは破棄して頂けるようお願いいたします<br><br><br>このメールへの返信はできません。"
	content := mail.Mail{
		Subject:     "認証コードのご連絡",
		To:          r.Mail,
		TextContent: text,
		HtmlContent: html,
	}
	//メール送信
	err = mail.SendMail(content)
	if err != nil {
		return "", nil
	}

	return userID.String(), nil
}

func (a *AuthUseCase) Login(r request.SignInRequest) (*entity.User, string, error) {
	user, err := a.authRepository.GetPassword(r.Mail)
	if err != nil {
		return nil, "", err
	}
	err = utils.CompareHashAndPassword(user.Password, r.Password)
	if err != nil {
		return nil, "", errors.New("not match password")
	}

	jwt, err := auth.IssueUserToken(user.UserID)
	if err != nil {
		return nil, "", errors.New("failed to generate JWT token")
	}
	return &user, jwt, nil
}

func (a *AuthUseCase) LoginMobile(r request.SignInRequest) (*entity.User, string, error) {
	user, err := a.authRepository.GetPassword(r.Mail)
	if err != nil {
		return nil, "", err
	}
	err = utils.CompareHashAndPassword(user.Password, r.Password)
	if err != nil {
		return nil, "", errors.New("not match password")
	}

	jwt, err := auth.IssueMobileUserToken(user.UserID)
	if err != nil {
		return nil, "", errors.New("failed to generate JWT token")
	}
	return &user, jwt, nil
}

func (a *AuthUseCase) CheckMail(r request.AuthCodeRequest) error {

	code, err := a.authRepository.CheckMailAddr(r.UserID)
	if err != nil {
		return err
	}
	if code != r.Code {
		return errors.New("not match code")
	}
	err = a.authRepository.UpdateStatus(r.UserID)
	if err != nil {
		return err
	}

	return nil
}
