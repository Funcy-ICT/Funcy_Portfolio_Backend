package infrastructure

import (
	"backend/app/domain/entity"
	"backend/app/domain/repository"

	"context"
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type userRepositoryImpl struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.AuthRepository {
	return &userRepositoryImpl{db: db}
}

// VALUES("5", $2, $3, $4, $5, $6, $7, $8, $9, $10)
func (ur *userRepositoryImpl) InsertAccount(user *entity.User) error {
	_, err := ur.db.Exec(`INSERT INTO users (id,display_name,icon,family_name,first_name,mail,password,grade,course,token,code,status)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		user.UserID, user.DisplayName, user.Icon, user.FamilyName, user.FirstName, user.Mail, user.Password, user.Grade, user.Course, user.Token, user.AuthCode, user.Status)
	if err != nil {
		log.Println(err)
		return errors.Wrap(err, "failed to insert")
	}

	_, err = func() (sql.Result, error) {
		return ur.db.DB.ExecContext(context.Background(), `INSERT INTO user_profile (user_id,header_image,bio) VALUES(?,?,?)`, user.UserID, "", "")
	}()
	if err != nil {
		return errors.Wrap(err, "failed to insert")
	}
	return nil
}

func (ur *userRepositoryImpl) CheckMailAddr(userID string) (string, error) {
	var user entity.User
	err := ur.db.Get(&user, `SELECT code FROM users WHERE id=?`, userID)
	if err != nil {
		return "", errors.New("Not a valid code")
	}
	return user.AuthCode, nil
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
	err := ur.db.Get(&user, `SELECT id,password,token FROM users WHERE mail=?`, mail)
	if err != nil {
		return user, errors.Wrap(err, "failed to find mail")
	}
	return user, nil
}
