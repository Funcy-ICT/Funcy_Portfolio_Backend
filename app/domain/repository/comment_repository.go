package repository

import "backend/app/domain/entity"

type CommentRepository interface {
	SelectCommentsByWorksID(worksID string) ([]*entity.Comment, error)
	InsertComment(comment *entity.Comment) error
}
