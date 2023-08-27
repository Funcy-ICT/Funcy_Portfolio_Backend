package infrastructure

import (
	"backend/app/domain/entity"
	"backend/app/domain/repository"
	"log"

	"github.com/jmoiron/sqlx"
)

type workRepositoryImpl struct {
	db *sqlx.DB
}

func NewWorkRepository(db *sqlx.DB) repository.WorkRepository {
	return &workRepositoryImpl{db: db}
}

func (ur *workRepositoryImpl) InsertWork(userID string, work *entity.WorkTable, images *[]entity.Image, tags *[]entity.Tag) error {
	log.Println("nannde")
	tx, _ := ur.db.Beginx()

	_, err := tx.Exec(`INSERT INTO works (id,user_id,title,description,thumbnail,url,movie_url,security) 
VALUES (?,?,?,?,?,?,?,?)`,
		work.ID, userID, work.Title, work.Description, work.Thumbnail, work.WorkUrl, work.MovieUrl, work.Security)
	if err != nil {
		return err
	}

	_, err = tx.NamedExec(`INSERT INTO work_images (id,work_id,image_url) VALUES (:id,:work_id,:image_url)`,
		*images)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.NamedExec("INSERT INTO `work_tags` (id,work_id,tag) VALUES (:id,:work_id,:tag)",
		*tags)
	if err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (ur *workRepositoryImpl) SelectWorks(numberOfWorks uint) (*[]*entity.ReadWorksList, error) {
	works := new([]*entity.ReadWorksList)
	err := ur.db.Select(
		works,
		"SELECT works.id, works.user_id, works.title, works.description, works.thumbnail, users.icon FROM works INNER JOIN users ON works.user_id = users.id ORDER BY works.created_at DESC LIMIT ?",
		numberOfWorks)
	if err != nil {
		return nil, err
	}

	return works, nil
}

func (ur *workRepositoryImpl) SelectWorkUser(userID string) (*entity.User, error) {
	var user entity.User
	err := ur.db.Get(
		&user,
		"select users.icon, users.display_name from users where users.id = ?",
		userID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *workRepositoryImpl) SelectWorksByUserID(userID string) (*[]*entity.ReadWorksList, error) {
	works := new([]*entity.ReadWorksList)
	err := ur.db.Select(
		works,
		"SELECT works.id, works.title, works.thumbnail, works.description, users.icon FROM works INNER JOIN users ON works.user_id = users.id WHERE works.user_id = ?;",
		userID)
	if err != nil {
		return nil, err
	}

	return works, nil
}

func (ur *workRepositoryImpl) SelectWork(workID string) (*entity.ReadWork, error) {
	work := new(entity.ReadWork)

	if err := ur.db.Get(work,
		"SELECT works.title, works.description, works.user_id, works.thumbnail, works.url, works.movie_url, works.security from works where works.id = ?", workID); err != nil {
		return nil, err
	}

	if err := ur.db.Select(&work.ImageURLs,
		"SELECT work_images.image_url from work_images where work_images.work_id = ?", workID); err != nil {
		log.Println("urls", err)
		return nil, err
	}

	if err := ur.db.Select(&work.Tags,
		"SELECT work_tags.tag from work_tags where work_tags.work_id = ?", workID); err != nil {
		log.Println("tags", err)
		return nil, err
	}

	return work, nil
}

func (ur *workRepositoryImpl) DeleteWork(workID string) error {
	tx, _ := ur.db.Beginx()

	_, err := tx.Exec("DELETE FROM funcy.works WHERE id=?", workID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM funcy.work_tags WHERE work_id=?", workID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM funcy.work_images WHERE work_id=?", workID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (ur *workRepositoryImpl) UpdateWork(work *entity.WorkTable, images *[]entity.Image, tags *[]entity.Tag) error {
	tx, _ := ur.db.Beginx()

	_, err := tx.Exec("UPDATE funcy.works SET title=?, description=?, thumbnail=?, url=?, movie_url=?, security=? WHERE id=?", work.Title, work.Description, work.Thumbnail, work.WorkUrl, work.MovieUrl, work.Security, work.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM funcy.work_images WHERE work_id=?", work.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM funcy.work_tags WHERE work_id=?", work.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.NamedExec(`INSERT INTO work_images (id,work_id,image_url) VALUES (:id,:work_id,:image_url)`,
		*images)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.NamedExec("INSERT INTO work_tags (id,work_id,tag) VALUES (:id,:work_id,:tag)",
		*tags)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
