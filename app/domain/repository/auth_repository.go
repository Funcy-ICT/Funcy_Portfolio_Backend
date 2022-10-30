package repository

import "backend/app/domain/entity"

type AuthRepository interface {
	InsertAccount(user *entity.User) error
}
