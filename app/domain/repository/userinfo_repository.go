package repository

import "backend/app/domain/entity"

type UserinfoRepository interface {
	SelectUserinfoByUserID(userID string) (*entity.Userinfo, error)
}
