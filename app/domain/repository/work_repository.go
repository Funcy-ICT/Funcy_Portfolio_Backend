package repository

import "backend/app/domain/entity"

type WorkRepository interface {
	InsertWork(userID string, work *entity.InsertWork, images *[]entity.Image, tags *[]entity.Tag) error
	SelectWork(workID string) (*entity.ReadWork, error)
	SelectWorkUser(userID string) (*entity.User, error)
	SelectWorks(numberOfWorks uint) (*[]*entity.ReadWorksList, error)
	SelectWorksByUserID(userID string) (*[]*entity.ReadWorksList, error)
	DeleteWork(workID string) error
	UpdateWork(work *entity.UpdateWork, images *[]entity.Image, tags *[]entity.Tag) error
}
