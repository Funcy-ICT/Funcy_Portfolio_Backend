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

func (ur *workRepositoryImpl) ReadWork(workID string) (*entity.ReadWork, error) {
	tx, _ := ur.db.Beginx()

	work := new(entity.ReadWork)
	if err := tx.Get(work, "SELECT title,description,url,movie_url,security from works where id = ?", workID); err != nil {
		return nil, err
	}
	if err := tx.Select(&work.Tags, "SELECT * from work_tags where work_id = ?", workID); err != nil {
		return nil, err
	}
	if err := tx.Select(&work.Images, "SELECT * from work_images where work_id = ?", workID); err != nil {
		return nil, err
	}

	return work, nil
}
