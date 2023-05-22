package infrastructure

import (
	"backend/app/domain/entity"
	"backend/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type userinfoRepositoryImpl struct {
	db *sqlx.DB
}

func NewUserInfoRepository(db *sqlx.DB) repository.UserinfoRepository {
	return &userinfoRepositoryImpl{db: db}
}

func (ur *userinfoRepositoryImpl) SelectUserinfoByUserID(userID string) (*entity.Userinfo, error) {
	// select Profile
	profile := new(entity.Profile)
	{
		err := ur.db.Get(
			profile,
			"SELECT user_id, header_image, bio FROM user_profile WHERE user_id = ? LIMIT 1;",
			userID)
		if err != nil {
			return nil, err
		}
	}

	// select Groups
	groups := new([]*entity.Group)
	{
		err := ur.db.Select(
			groups,
			"SELECT id, user_id, role, status FROM groups WHERE user_id = ?;",
			userID)
		if err != nil {
			return nil, err
		}
	}

	// select skills
	skills := new([]*entity.Skill)
	{
		err := ur.db.Select(
			skills,
			"SELECT skill_name, user_id FROM skills WHERE user_id = ?;",
			userID)
		if err != nil {
			return nil, err
		}
	}

	// select sns
	sns := new([]*entity.SNS)
	{
		err := ur.db.Select(
			sns,
			"SELECT user_id, sns Where user_id = ? FROM sns;",
			userID)
		if err != nil {
			return nil, err
		}
	}

	return &entity.Userinfo{
		Profile: profile,
		Groups:  groups,
		Skills:  skills,
		SNS:     sns,
	}, nil
}
