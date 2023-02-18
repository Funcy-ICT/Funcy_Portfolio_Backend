package repository

import "backend/app/domain/entity"

type AuthRepository interface {
	InsertAccount(user *entity.User) error
	GetPassword(mail string) (entity.User, error)
	CheckMailAddr(userID string) (string, error)
	UpdateStatus(userID string) error
	GetToken(userID string) (string, error)
}
