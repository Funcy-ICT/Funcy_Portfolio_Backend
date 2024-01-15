package repository

import "backend/app/domain/entity"

type UserinfoRepository interface {
	CreateNewUserinfo(userinfo *entity.Userinfo) error
	SelectUserinfoByUserID(userID string) (*entity.Userinfo, error)
	UpdateUserinfo(userinfo *entity.UpdateUserinfo) error
}
