package repository

import "backend/app/domain/entity"

type WorkRepository interface {
	InsertWork(userID string, work *entity.WorkTable, images *[]entity.Image, tags *[]entity.Tag) error
	SelectWork(workID string) (*entity.ReadWork, error)
	SelectWorks(numberOfWorks uint) (*[]*entity.ReadWorksList, error)
	SelectWorksByUserID(userID string) (*[]*entity.ReadWorksList, error)
	DeleteWork(workID string) error
}
