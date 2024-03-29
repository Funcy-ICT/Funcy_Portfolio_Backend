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
			"SELECT UP.user_id, UP.header_image, UP.bio, U.display_name, U.icon "+
				"FROM user_profile AS UP "+
				"INNER JOIN users AS U "+
				"ON UP.user_id = U.id "+
				"WHERE UP.user_id = ? LIMIT 1;",
			userID)
		if err != nil {
			return nil, err
		}
	}

	// select Groups
	groups := new([]*entity.GroupMember)
	{
		err := ur.db.Select(
			groups,
			"SELECT GM.group_id, G.name, GM.user_id, GM.role, GM.status "+
				"FROM group_member AS GM "+
				"INNER JOIN groups AS G "+
				"ON GM.group_id = G.id "+
				"WHERE user_id = ?;",
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
			"SELECT user_id, sns FROM sns WHERE user_id = ?;",
			userID)
		if err != nil {
			return nil, err
		}
	}

	return &entity.Userinfo{
		Profile:      profile,
		JoinedGroups: groups,
		Skills:       skills,
		SNS:          sns,
	}, nil
}

func (ur *userinfoRepositoryImpl) CreateNewUserinfo(userinfo *entity.Userinfo) error {
	tx, err := ur.db.Beginx()
	if err != nil {
		return err
	}

	{
		// user profile
		_, err := tx.NamedExec(
			"INSERT INTO user_profile (user_id, header_image, bio) VALUES (:user_id, :header_image, :bio);",
			userinfo.Profile,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// group member
	for _, group := range *userinfo.JoinedGroups {
		_, err := tx.NamedExec(
			"INSERT INTO group_member (group_id, user_id, role, status) VALUES (:group_id, :user_id, :role, :status);",
			group,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// skills
	for _, skill := range *userinfo.Skills {
		_, err := tx.NamedExec(
			"INSERT INTO skills (skill_name, user_id) VALUES (:skill_name, :user_id);",
			skill,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// sns
	for _, sns := range *userinfo.SNS {
		_, err := tx.NamedExec(
			"INSERT INTO sns (user_id, sns) VALUES (:user_id, :sns);",
			sns,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
