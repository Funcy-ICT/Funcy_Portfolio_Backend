package usecase

import (
	"backend/app/domain/entity"
	"backend/app/domain/repository"
	"backend/app/interfaces/request"
	"backend/app/interfaces/response"

	"errors"
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

func (w *WorkUseCase) ReadWorks(numberOfWorks uint) (*[]*entity.ReadWorksList, error) {
	works, err := w.workRepository.SelectWorks(numberOfWorks)
	if err != nil {
		return nil, err
	}
	return works, nil
}

func (w *WorkUseCase) ReadWork(workID string) (*entity.ReadWork, error) {
	work, err := w.workRepository.SelectWork(workID)
	if err != nil {
		return nil, err
	}
	return work, nil
}

func (w *WorkUseCase) DeleteWork(workID string) error {
	_, err := w.workRepository.SelectWork(workID)
	if err != nil {
		return errors.New(response.NoRows)
	}

	err = w.workRepository.DeleteWork(workID)
	if err != nil {
		return err
	}
	return nil
}
