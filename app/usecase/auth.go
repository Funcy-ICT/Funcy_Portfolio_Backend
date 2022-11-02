package usecase

import (
	"backend/app/domain/entity"
	"backend/app/domain/repository"
	"backend/app/interfaces/request"
	"backend/app/packages/utils"
	"backend/app/packages/utils/auth"

	"errors"

	"github.com/google/uuid"
	"github.com/mileusna/useragent"
)

type AuthUseCase struct {
	authRepository repository.AuthRepository
}

func NewAuthUseCase(authRepository repository.AuthRepository) *AuthUseCase {
	return &AuthUseCase{authRepository: authRepository}
}

func (a *AuthUseCase) CreateAccount(r request.SignUpRequest) error {

	userID, err := uuid.NewRandom()
	if err != nil {
		return errors.New("userID generate is failed")
	}
	r.Password, err = utils.PasswordEncrypt(r.Password)
	if err != nil {
		return errors.New("password generate is failed")
	}
	token, err := uuid.NewRandom()
	if err != nil {
		return errors.New("tokenID generate is failed")
	}

	user, err := entity.NewUser(&r, userID.String(), token.String())
	if err != nil {
		return err
	}

	err = a.authRepository.InsertAccount(user)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthUseCase) Login(r request.SignInRequest, agent string) (string, error) {
	user, err := a.authRepository.GetPassword(r.Mail)
	if err != nil {
		return "", err
	}
	err = utils.CompareHashAndPassword(user.Password, r.Password)
	if err != nil {
		return "", errors.New("not match password")
	}

	ua := useragent.Parse(agent)
	switch {
	case ua.Mobile == true:
		jwt, _ := auth.IssueMobileUserToken(user.UserID)
		return jwt, nil
	case ua.Desktop == true || ua.Name == "PostmanRuntime":
		jwt, _ := auth.IssueUserToken(user.UserID)
		return jwt, nil
	default:
		return "", errors.New("not permitted os")
	}
}
