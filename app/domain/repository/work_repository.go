package repository

import "backend/app/domain/entity"

type WorkRepository interface {
	InsertWork(userID string, work *entity.WorkTable, images *[]entity.Image, tags *[]entity.Tag) error
	ReadWork(workID string) (*entity.ReadWork, error)
}
