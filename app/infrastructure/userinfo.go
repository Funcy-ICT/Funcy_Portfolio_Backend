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
	// select Profile and Course
	var profileWithCourse struct {
		entity.Profile
		Course string `db:"course"`
	}
	{
		err := ur.db.Get(
			&profileWithCourse,
			"SELECT UP.user_id, UP.header_image, UP.bio, U.display_name, U.icon, U.course, U.mail "+
				"FROM user_profile AS UP "+
				"INNER JOIN users AS U "+
				"ON UP.user_id = U.id "+
				"WHERE UP.user_id = ? LIMIT 1;",
			userID)
		if err != nil {
			return nil, err
		}
	}

	profile := &profileWithCourse.Profile

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
		Course:       profileWithCourse.Course,
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

func (ur *userinfoRepositoryImpl) UpdateUserinfo(userinfo *entity.UpdateUserinfo) error {
	tx, err := ur.db.Beginx()
	if err != nil {
		return err
	}

	// Update
	{
		// user profile
		{
			_, err := tx.NamedExec(
				"UPDATE user_profile SET header_image=:header_image, bio=:bio WHERE user_id=:user_id;",
				userinfo.Profile,
			)
			if err != nil {
				tx.Rollback()
				return err
			}
		}

		// user icon
		{
			_, err := tx.NamedExec(
				"UPDATE users SET icon=:icon, display_name=:display_name WHERE id=:user_id;",
				map[string]interface{}{
					"icon":         userinfo.Profile.Icon,
					"display_name": userinfo.Profile.DisplayName,
					"user_id":      userinfo.Profile.UserID,
				},
			)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	// Delete
	{
		// sns
		{
			_, err := tx.Exec(
				"DELETE FROM sns WHERE user_id = ?",
				userinfo.Profile.UserID,
			)
			if err != nil {
				tx.Rollback()
				return err
			}
		}

		// skills
		{
			_, err := tx.Exec(
				"DELETE FROM skills WHERE user_id = ?",
				userinfo.Profile.UserID,
			)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	// Re-Create
	{
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
	}

	return tx.Commit()
}

func (ur *userinfoRepositoryImpl) SearchUsersByKeyword(keyword string, limit uint) (*[]entity.UserSearchResult, error) {
	searchPattern := "%" + keyword + "%"

	query := `
		SELECT DISTINCT u.id as user_id, u.display_name, u.icon, u.course
		FROM users u
		LEFT JOIN user_profile up ON u.id = up.user_id
		LEFT JOIN skills s ON u.id = s.user_id
		WHERE u.display_name LIKE ?
			OR u.course LIKE ?
			OR up.bio LIKE ?
			OR s.skill_name LIKE ?
		LIMIT ?
	`

	type UserRow struct {
		UserID      string `db:"user_id"`
		DisplayName string `db:"display_name"`
		Icon        string `db:"icon"`
		Course      string `db:"course"`
	}

	var userRows []UserRow
	err := ur.db.Select(&userRows, query, searchPattern, searchPattern, searchPattern, searchPattern, limit)
	if err != nil {
		return nil, err
	}

	// 各ユーザーのスキルを取得
	results := make([]entity.UserSearchResult, 0, len(userRows))
	for _, row := range userRows {
		skills := []string{}
		err := ur.db.Select(&skills, "SELECT skill_name FROM skills WHERE user_id = ?", row.UserID)
		if err != nil {
			// スキルがない場合はエラーを無視
			skills = []string{}
		}

		results = append(results, entity.UserSearchResult{
			UserID:      row.UserID,
			DisplayName: row.DisplayName,
			Icon:        row.Icon,
			Course:      row.Course,
			Skills:      skills,
		})
	}

	return &results, nil
}
