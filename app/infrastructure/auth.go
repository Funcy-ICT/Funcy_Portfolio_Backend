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

//VALUES("5", $2, $3, $4, $5, $6, $7, $8, $9, $10)
func (ur *userRepositoryImpl) InsertAccount(user *entity.User) error {
	_, err := ur.db.Exec(`INSERT INTO users (id,display_name,icon,family_name,first_name,mail,password,grade,course,token)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		user.UserID, user.DisplayName, user.Icon, user.FamilyName, user.FirstName, user.Mail, user.Password, user.Grade, user.Course, user.Token)
	if err != nil {
		log.Println(err)
		return errors.Wrap(err, "failed to insert")
	}
	return nil
}

//func (ur *userRepositoryImpl) CheckMailAddr(mail string) error {
//	var user entity.User
//	err := ur.db.Get(&user, `SELECT mail FROM users WHERE mail=?`, mail)
//	if err != nil {
//		return err
//	}
//	if user.Mail == "" {
//		return nil
//	}
//	return errors.New("exist mail address")
//}

func (ur *userRepositoryImpl) GetPassword(mail string) (entity.User, error) {
	var user entity.User
	err := ur.db.Get(&user, `SELECT id,password,token FROM users WHERE mail=?`, mail)
	if err != nil {
		return user, errors.Wrap(err, "failed to find mail")
	}
	return user, nil
}
