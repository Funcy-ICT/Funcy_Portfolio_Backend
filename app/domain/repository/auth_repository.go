package repository

import "backend/app/domain/entity"

type AuthRepository interface {
	InsertAccount(user *entity.User) error
	GetPassword(mail string) (entity.User, error)
	CheckMailAddr(userID string) (string, error)
	UpdateStatus(userID string) error
	GetByAuth0Sub(auth0Sub string) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	CreateFromAuth0(user *entity.User) error
	UpdateUser(user *entity.User) error
}
