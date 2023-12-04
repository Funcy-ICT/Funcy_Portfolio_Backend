package usecase

import (
	"backend/app/domain/entity"
	"backend/app/domain/repository"

	"github.com/google/uuid"
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
		image := entity.Image{
			ID:     uuid.NewString(),
			WorkID: work.ID,
			Image:  v,
		}
		imagesEntity = append(imagesEntity, image)
	}

	tagsEntity := make([]entity.Tag, 0, len(tags))
	for _, v := range tags {
		tag := entity.Tag{
			ID:     uuid.NewString(),
			WorkID: work.ID,
			Tag:    v,
		}
		tagsEntity = append(tagsEntity, tag)
	}

	err := w.workRepository.InsertWork(userId, work, &imagesEntity, &tagsEntity)
	if err != nil {
		return "", err
	}

	return work.ID, nil
}

func (w *WorkUseCase) ReadWorks(numberOfWorks uint) (*[]*entity.ReadWorksList, error) {
	works, err := w.workRepository.SelectWorks(numberOfWorks)
	if err != nil {
		return &[]*entity.ReadWorksList{}, err
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
		// no rows set
		// ここでは、エラーを返さない
		return nil
	}

	err = w.workRepository.DeleteWork(workID)
	if err != nil {
		return err
	}
	return nil
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

	work := &entity.UpdateWork{
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
		image := entity.Image{
			ID:     uuid.NewString(),
			WorkID: work.ID,
			Image:  v,
		}
		imagesEntity = append(imagesEntity, image)
	}

	tagsEntity := make([]entity.Tag, 0, len(tags))
	for _, v := range tags {
		tag := entity.Tag{
			ID:     uuid.NewString(),
			WorkID: work.ID,
			Tag:    v,
		}
		tagsEntity = append(tagsEntity, tag)
	}

	err := w.workRepository.UpdateWork(work, &imagesEntity, &tagsEntity)
	if err != nil {
		return err
	}

	return nil
}
