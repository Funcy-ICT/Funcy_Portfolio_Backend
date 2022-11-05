package usecase

import (
	"backend/app/domain/entity"
	"backend/app/domain/repository"
	"backend/app/interfaces/request"
)

type WorkUseCase struct {
	workRepository repository.WorkRepository
}

func NewWorkUseCase(workRepository repository.WorkRepository) *WorkUseCase {
	return &WorkUseCase{workRepository: workRepository}
}

func (w *WorkUseCase) CreateWork(r request.CreateWorkRequest, userId string) (string, error) {

	work, err := entity.NewWork(r)
	if err != nil {
		return "", err
	}
	images, err := entity.NewWorkImages(r, work.ID)
	if err != nil {
		return "", err
	}
	tags, err := entity.NewWorkTags(r, work.ID)
	if err != nil {
		return "", err
	}

	err = w.workRepository.InsertWork(userId, work, images, tags)
	if err != nil {
		return "", err
	}
	return work.ID, nil
}
