package usecase

import (
	"backend/app/domain/entity"
	"backend/app/domain/repository"
	"time"

	"github.com/google/uuid"
)

type CommentUseCase struct {
	commentRepository repository.CommentRepository
}

func NewCommentUseCase(
	commentRepository repository.CommentRepository,
) *CommentUseCase {
	return &CommentUseCase{
		commentRepository: commentRepository,
	}
}

func (u *CommentUseCase) GetComment(workID string) ([]*entity.Comment, error) {
	return u.commentRepository.SelectCommentsByWorksID(workID)
}

func (u *CommentUseCase) CreateComment(userID, worksID, content string) (string, error) {
	comment := &entity.Comment{
		ID:        uuid.NewString(),
		UserID:    userID,
		WorksID:   worksID,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := u.commentRepository.InsertComment(comment)
	if err != nil {
		return "", err
	}
	return comment.ID, nil
}
