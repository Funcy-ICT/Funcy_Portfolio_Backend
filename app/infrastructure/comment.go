package infrastructure

import (
	"backend/app/domain/entity"
	"time"

	"github.com/jmoiron/sqlx"
)

type CommentRepository interface {
	SelectCommentsByWorksID(worksID string) ([]*entity.Comment, error)
}

type commentRepositoryImpl struct {
	db *sqlx.DB
}

func NewCommentRepository(db *sqlx.DB) CommentRepository {
	return &commentRepositoryImpl{db: db}
}

func (ur *commentRepositoryImpl) SelectCommentsByWorksID(worksID string) ([]*entity.Comment, error) {
	var rows *sqlx.Rows
	var err error
	rows, err = ur.db.Queryx("SELECT id, user_id, works_id, text, created_at, updated_at FROM comment WHERE works_id = ?", worksID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// time.time型への変換
	var comments []*entity.Comment
	for rows.Next() {
		var comment entity.Comment
		var createdAt, updatedAt []byte
		if err := rows.Scan(&comment.ID, &comment.UserID, &comment.WorksID, &comment.Text, &createdAt, &updatedAt); err != nil {
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
