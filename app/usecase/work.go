package usecase

import (
	"backend/app/domain/entity"
	"backend/app/domain/repository"
	"errors"
)

type WorkUseCase struct {
	workRepository repository.WorkRepository
}

func NewWorkUseCase(workRepository repository.WorkRepository) *WorkUseCase {
	return &WorkUseCase{workRepository: workRepository}
}

func (w *WorkUseCase) CreateWork(
	userId string,
	title string,
	description string,
	thumbnail string,
	workUrl string,
	movieUrl string,
	groupID string,
	security int,
	images []string,
	tags []string,
) (string, error) {

	work := entity.NewInsertWork(
		title,
		description,
		thumbnail,
		workUrl,
		movieUrl,
		groupID,
		security,
	)

	imagesEntity := make([]entity.Image, 0, len(images))
	for _, v := range images {
		imagesEntity = append(imagesEntity, entity.NewImage(work.ID, v))
	}

	tagsEntity := make([]entity.Tag, 0, len(tags))
	for _, v := range tags {
		tagsEntity = append(tagsEntity, entity.NewTag(work.ID, v))
	}

	err := w.workRepository.InsertWork(userId, &work, &imagesEntity, &tagsEntity)
	if err != nil {
		return "", err
	}

	return work.ID, nil
}

func (w *WorkUseCase) ReadWorks(numberOfWorks uint, tag string) (*[]*entity.ReadWorksList, error) {
	if len(tag) == 0 {
		works, err := w.workRepository.SelectWorks(numberOfWorks)
		if err != nil {
			return &[]*entity.ReadWorksList{}, err
		}

		return works, nil
	} else {
		works, err := w.workRepository.SelectWorksByTag(numberOfWorks, tag)
		if err != nil {
			return &[]*entity.ReadWorksList{}, err
		}

		return works, nil
	}
}

func (w *WorkUseCase) ReadWork(workID string) (*entity.ReadWork, *entity.User, error) {
	work, err := w.workRepository.SelectWork(workID)
	if err != nil {
		return nil, nil, errors.New("failed to retrieve work")
	}

	user, err := w.workRepository.SelectWorkUser(work.UserId)
	if err != nil {
		return nil, nil, errors.New("failed to retrieve work user")
	}

	return work, user, nil
}

func (w *WorkUseCase) DeleteWork(workID string) error {
	return w.workRepository.DeleteWork(workID)
}

func (w *WorkUseCase) UpdateWork(
	workID string,
	title string,
	description string,
	thumbnail string,
	workUrl string,
	movieUrl string,
	groupID string,
	security int,
	images []string,
	tags []string) error {

	work := entity.UpdateWork{
		ID:          workID,
		Title:       title,
		Description: description,
		Thumbnail:   thumbnail,
		WorkUrl:     workUrl,
		MovieUrl:    movieUrl,
		GroupID:     groupID,
		Security:    security,
	}

	imagesEntity := make([]entity.Image, 0, len(images))
	for _, v := range images {
		imagesEntity = append(imagesEntity, entity.NewImage(work.ID, v))
	}

	tagsEntity := make([]entity.Tag, 0, len(tags))
	for _, v := range tags {
		tagsEntity = append(tagsEntity, entity.NewTag(work.ID, v))
	}

	err := w.workRepository.UpdateWork(&work, &imagesEntity, &tagsEntity)
	if err != nil {
		return err
	}

	return nil
}
