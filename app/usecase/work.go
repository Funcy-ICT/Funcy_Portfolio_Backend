package usecase

import (
	"backend/app/domain/entity"
	"backend/app/domain/repository"
	"backend/app/interfaces/request"
	"backend/app/interfaces/response"

	"errors"

	"github.com/google/uuid"
)

type WorkUseCase struct {
	workRepository repository.WorkRepository
}

func NewWorkUseCase(workRepository repository.WorkRepository) *WorkUseCase {
	return &WorkUseCase{workRepository: workRepository}
}

func (w *WorkUseCase) CreateWork(r request.CreateWorkRequest, userId string) (string, error) {

	workId := uuid.NewString()

	work := &entity.WorkTable{
		ID:          workId,
		Title:       r.Title,
		Description: r.Description,
		URL:         r.WorkUrl,
		MovieUrl:    r.MovieUrl,
		Security:    r.Security,
	}

	images := make([]entity.Image, 0, len(r.Images))
	for _, v := range r.Images {
		image := entity.Image{
			ID:     uuid.NewString(),
			WorkID: workId,
			Image:  v.Image,
		}
		images = append(images, image)
	}

	tags := make([]entity.Tag, 0, len(r.Tags))
	for _, v := range r.Tags {
		tag := entity.Tag{
			ID:     uuid.NewString(),
			WorkID: workId,
			Tag:    v.Tag,
		}
		tags = append(tags, tag)
	}

	err := w.workRepository.InsertWork(userId, work, &images, &tags)
	if err != nil {
		return "", err
	}

	return workId, nil
}

func (w *WorkUseCase) ReadWorks(numberOfWorks uint) (*[]*entity.ReadWorksList, error) {
	works, err := w.workRepository.SelectWorks(numberOfWorks)
	if err != nil {
		return nil, err
	}
	return works, nil
}

func (w *WorkUseCase) ReadWork(workID string) (*entity.ReadWork, *entity.User, error) {
	work, err := w.workRepository.SelectWork(workID)
	if err != nil {
		return nil, nil, err
	}
	user, err := w.workRepository.SelectWorkUser(work.UserId)
	if err != nil {
		return nil, nil, err
	}
	return work, user, nil
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
