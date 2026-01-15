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

func (ur *workRepositoryImpl) InsertWork(userID string, work *entity.InsertWork, images *[]entity.Image, tags *[]entity.Tag) error {
	tx, err := ur.db.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`INSERT INTO works (id, user_id, title, description, thumbnail, url, movie_url, security, group_id) 
VALUES (?,?,?,?,?,?,?,?,?)`,
		work.ID, userID, work.Title, work.Description, work.Thumbnail, work.WorkUrl, work.MovieUrl, work.Security, work.GroupID)
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
		"SELECT works.id, works.title, works.description, works.thumbnail, works.security, users.icon, works.user_id FROM works INNER JOIN users ON works.user_id = users.id WHERE works.security = 1 ORDER BY works.created_at DESC LIMIT ?",
		numberOfWorks)
	if err != nil {
		return nil, err
	}

	return works, nil
}

func (ur *workRepositoryImpl) SelectWorksByTag(numberOfWorks uint, tag string) (*[]*entity.ReadWorksList, error) {
	works := new([]*entity.ReadWorksList)
	err := ur.db.Select(
		works,
		"SELECT works.id, works.title, works.description, works.thumbnail, works.security, users.icon, works.user_id FROM works INNER JOIN work_images ON works.id = work_images.work_id INNER JOIN work_tags ON works.id = work_tags.work_id INNER JOIN users ON works.user_id = users.id WHERE work_tags.tag=? ORDER BY works.created_at DESC LIMIT ?",
		tag,
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
		"SELECT works.id, works.title, works.thumbnail, works.description, works.security, users.icon, works.user_id FROM works INNER JOIN users ON works.user_id = users.id WHERE works.user_id = ? ORDER BY works.created_at DESC;",
		userID)
	if err != nil {
		return nil, err
	}

	return works, nil
}

func (ur *workRepositoryImpl) SearchWorksByKeyword(keyword string, limit uint, scope string) (*[]*entity.ReadWorksList, error) {
	works := new([]*entity.ReadWorksList)
	searchPattern := "%" + keyword + "%"

	var query string
	var args []interface{}

	switch scope {
	case "tag":
		query = `SELECT DISTINCT works.id, works.title, works.description, works.thumbnail, works.security, users.icon, works.user_id, works.created_at
				FROM works
				INNER JOIN users ON works.user_id = users.id
				LEFT JOIN work_tags ON works.id = work_tags.work_id
				WHERE works.security = 1 AND work_tags.tag LIKE ?
				ORDER BY works.created_at DESC LIMIT ?`
		args = []interface{}{searchPattern, limit}
	case "title":
		query = `SELECT DISTINCT works.id, works.title, works.description, works.thumbnail, works.security, users.icon, works.user_id, works.created_at
				FROM works
				INNER JOIN users ON works.user_id = users.id
				WHERE works.security = 1 AND works.title LIKE ?
				ORDER BY works.created_at DESC LIMIT ?`
		args = []interface{}{searchPattern, limit}
	default: // "all"
		query = `SELECT DISTINCT works.id, works.title, works.description, works.thumbnail, works.security, users.icon, works.user_id, works.created_at
				FROM works
				INNER JOIN users ON works.user_id = users.id
				LEFT JOIN work_tags ON works.id = work_tags.work_id
				WHERE works.security = 1 AND (works.title LIKE ? OR works.description LIKE ? OR (work_tags.tag IS NOT NULL AND work_tags.tag LIKE ?))
				ORDER BY works.created_at DESC LIMIT ?`
		args = []interface{}{searchPattern, searchPattern, searchPattern, limit}
	}

	err := ur.db.Select(works, query, args...)
	if err != nil {
		log.Printf("SearchWorksByKeyword SQL error: %v, query: %s, args: %v", err, query, args)
		return nil, err
	}

	log.Printf("SearchWorksByKeyword found %d works", len(*works))
	return works, nil
}

func (ur *workRepositoryImpl) SelectWork(workID string) (*entity.ReadWork, error) {
	work := new(entity.ReadWork)

	if err := ur.db.Get(work,
		"SELECT works.title, works.description, works.user_id, works.thumbnail, works.url, works.movie_url, works.security, works.group_id from works where works.id = ?", workID); err != nil {
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
	tx, err := ur.db.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM funcy.works WHERE id=?", workID)
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

func (ur *workRepositoryImpl) UpdateWork(work *entity.UpdateWork, images *[]entity.Image, tags *[]entity.Tag) error {
	tx, err := ur.db.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE funcy.works SET title=?, description=?, thumbnail=?, url=?, movie_url=?, security=?, group_id=? WHERE id=?", work.Title, work.Description, work.Thumbnail, work.WorkUrl, work.MovieUrl, work.Security, work.GroupID, work.ID)
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
