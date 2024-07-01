package usecase

import (
	"backend/app/domain/entity"
	"backend/app/domain/repository"
)

type CommentUseCase struct {
	commentRepository repository.CommentRepository
}

func NewCommentUsecace(
	commentRepository repository.CommentRepository,
) *CommentUseCase {
	return &CommentUseCase{
		commentRepository: commentRepository,
	}
}

func (u *CommentUseCase) GetComment(workID string) ([]*entity.Comment, error) {
	return u.commentRepository.SelectCommentsByWorksID(workID)
}
