package infrastructure

import (
	"backend/app/domain/entity"
	"time"

	"github.com/jmoiron/sqlx"
)

type CommentRepository interface {
	SelectCommentsByWorksID(worksID string) ([]*entity.Comment, error)
	InsertComment(comment *entity.Comment) error
}

type commentRepositoryImpl struct {
	db *sqlx.DB
}

func NewCommentRepository(db *sqlx.DB) CommentRepository {
	return &commentRepositoryImpl{db: db}
}

func (ur *commentRepositoryImpl) SelectCommentsByWorksID(worksID string) ([]*entity.Comment, error) {
	query := `
		SELECT 
			c.id, c.user_id, c.works_id, c.content, c.created_at, c.updated_at,
			u.display_name, u.icon
		FROM 
			comment c
		JOIN 
			users u ON c.user_id = u.id
		WHERE 
			c.works_id = ?
	`

	rows, err := ur.db.Queryx(query, worksID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*entity.Comment
	for rows.Next() {
		var comment entity.Comment
		var createdAt, updatedAt []byte

		err := rows.Scan(
			&comment.ID, &comment.UserID, &comment.WorksID, &comment.Content, &createdAt, &updatedAt,
			&comment.CommentUserProfile.DisplayName, &comment.CommentUserProfile.Icon,
		)
		if err != nil {
			return nil, err
		}
		comment.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))
		if err != nil {
			return nil, err
		}
		comment.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", string(updatedAt))
		if err != nil {
			return nil, err
		}

		comments = append(comments, &comment)
	}

	return comments, nil
}

func (ur *commentRepositoryImpl) InsertComment(comment *entity.Comment) error {
	_, err := ur.db.Exec(`INSERT INTO comment (id, user_id, works_id, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`,
		comment.ID, comment.UserID, comment.WorksID, comment.Content, comment.CreatedAt, comment.UpdatedAt)
	return err
}
