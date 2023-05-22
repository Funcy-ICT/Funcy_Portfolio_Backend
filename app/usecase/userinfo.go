package usecase

import (
	"backend/app/domain/entity"
	"backend/app/domain/repository"
)

type UserinfoUseCase struct {
	userinfoRepository repository.UserinfoRepository
	workRepository     repository.WorkRepository
}

func NewUserinfoUsecace(
	userinfoRepository repository.UserinfoRepository,
	workRepository repository.WorkRepository,
) *UserinfoUseCase {
	return &UserinfoUseCase{
		userinfoRepository: userinfoRepository,
		workRepository:     workRepository,
	}
}

func (u *UserinfoUseCase) GetUserinfo(userID string) (*entity.Userinfo, error) {
	return u.userinfoRepository.SelectUserinfoByUserID(userID)
}
