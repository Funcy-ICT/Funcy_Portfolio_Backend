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

	_, err := tx.Exec(`INSERT INTO works (id,user_id,title,description,url,movie_url,security) 
VALUES (?,?,?,?,?,?,?)`,
		work.ID, userID, work.Title, work.Description, work.MovieUrl, work.URL, work.Security)
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
		"SELECT works.id, works.title, work_images.image_url, works.description, users.icon FROM works INNER JOIN work_images ON works.id = work_images.work_id INNER JOIN users ON works.user_id = users.id ORDER BY works.created_at DESC LIMIT ?",
		numberOfWorks)
	if err != nil {
		return nil, err
	}

	return works, nil
}
