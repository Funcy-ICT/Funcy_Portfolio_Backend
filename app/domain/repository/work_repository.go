package repository

import (
	"backend/app/domain/entity"
)

type WorkRepository interface {
	InsertWork(userID string, work *entity.WorkTable, images *[]entity.Image, tags *[]entity.Tag) error
	SelectWork(workID string) (*entity.ReadWork, error)
	SelectWorkUser(userID string) (*entity.User, error)
	SelectWorks(numberOfWorks uint) (*[]*entity.ReadWorksList, error)
	SelectWorksByTag(numberOfWorks uint, tag string) (*[]*entity.ReadWorksList, error)
	DeleteWork(workID string) error
	UpdateWork(work *entity.WorkTable, images *[]entity.Image, tags *[]entity.Tag) error
}
