package infrastructure

import (
	"backend/app/domain/entity"
	"backend/app/domain/repository"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"log"
)

type userRepositoryImpl struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.AuthRepository {
	return &userRepositoryImpl{db: db}
}

// VALUES("5", $2, $3, $4, $5, $6, $7, $8, $9, $10)
func (ur *userRepositoryImpl) InsertAccount(user *entity.User) error {
	tx, err := ur.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}

	_, err = tx.Exec(`INSERT INTO users (id,display_name,icon,family_name,first_name,mail,password,grade,course,token,code,status)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		user.UserID, user.DisplayName, user.Icon, user.FamilyName, user.FirstName, user.Mail, user.Password, user.Grade, user.Course, user.Token, user.AuthCode, user.Status)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return errors.Wrap(err, "failed to insert user")
	}

	_, err = tx.Exec(`INSERT INTO user_profile (user_id, header_image, bio) VALUES (?, ?, ?)`,
		user.UserID, "", "")
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return errors.Wrap(err, "failed to insert user_profile")
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}

	return nil
}

func (ur *userRepositoryImpl) CheckMailAddr(userID string) (string, error) {
	var user entity.User
	err := ur.db.Get(&user, `SELECT code FROM users WHERE id=?`, userID)
	if err != nil {
		return "", errors.New("Not a valid code")
	}
	if user.AuthCode == nil {
		return "", nil
	}
	return *user.AuthCode, nil
}

func (ur *userRepositoryImpl) UpdateStatus(userID string) error {

	_, err := ur.db.NamedExec(`UPDATE users SET status=:status where id=:userID`,
		map[string]interface{}{
			"status": "active",
			"userID": userID,
		})
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepositoryImpl) GetPassword(mail string) (entity.User, error) {
	var user entity.User
	err := ur.db.Get(&user, `SELECT id, auth0_sub, display_name, icon, family_name, first_name, mail, password, grade, course, token, status, code, created_at, updated_at FROM users WHERE mail=?`, mail)
	if err != nil {
		return user, errors.Wrap(err, "failed to find mail")
	}
	return user, nil
}

func (ur *userRepositoryImpl) GetByAuth0Sub(auth0Sub string) (*entity.User, error) {
	var user entity.User
	err := ur.db.Get(&user, `SELECT id, auth0_sub, display_name, icon, family_name, first_name, mail, password, grade, course, token, status, code, created_at, updated_at
		FROM users WHERE auth0_sub = ?`, auth0Sub)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepositoryImpl) CreateFromAuth0(user *entity.User) error {
	tx, err := ur.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}

	_, err = tx.Exec(`INSERT INTO users (id,display_name,icon,family_name,first_name,mail,grade,course,status,auth0_sub)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		user.UserID, user.DisplayName, user.Icon, user.FamilyName, user.FirstName, user.Mail, user.Grade, user.Course, user.Status, user.Auth0Sub)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "failed to insert user")
	}

	_, err = tx.Exec(`INSERT INTO user_profile (user_id, header_image, bio) VALUES (?, ?, ?)`,
		user.UserID, "", "")
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "failed to insert user_profile")
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}

	return nil
}

func (ur *userRepositoryImpl) UpdateUser(user *entity.User) error {
	_, err := ur.db.Exec(`UPDATE users SET display_name = ?, icon = ?, mail = ?, family_name = ?, first_name = ?, grade = ?, course = ?, status = ?, auth0_sub = ? WHERE id = ?`,
		user.DisplayName, user.Icon, user.Mail, user.FamilyName, user.FirstName, user.Grade, user.Course, user.Status, user.Auth0Sub, user.UserID)
	if err != nil {
		return errors.Wrap(err, "failed to update user")
	}
	return nil
}
