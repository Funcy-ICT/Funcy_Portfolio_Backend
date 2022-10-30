package usecase

import (
	"backend/app/domain/entity"
	"backend/app/domain/repository"
	"backend/app/interfaces/request"
	"backend/app/packages/utils"
	"errors"
	"github.com/google/uuid"
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
