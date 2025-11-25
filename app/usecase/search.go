package usecase

import (
	"backend/app/domain/entity"
	"backend/app/domain/repository"
)

type SearchUseCase struct {
	workRepository     repository.WorkRepository
	userinfoRepository repository.UserinfoRepository
}

func NewSearchUseCase(workRepository repository.WorkRepository, userinfoRepository repository.UserinfoRepository) *SearchUseCase {
	return &SearchUseCase{
		workRepository:     workRepository,
		userinfoRepository: userinfoRepository,
	}
}

func (uc *SearchUseCase) SearchWorks(keyword string, limit uint, scope string) (*[]*entity.ReadWorksList, error) {
	// デフォルト値の設定
	if limit == 0 {
		limit = 100
	}
	if scope == "" {
		scope = "all"
	}

	// 作品検索を実行
	works, err := uc.workRepository.SearchWorksByKeyword(keyword, limit, scope)
	if err != nil {
		return nil, err
	}

	return works, nil
}

func (uc *SearchUseCase) SearchUsers(keyword string, limit uint) (*[]entity.UserSearchResult, error) {
	// デフォルト値の設定
	if limit == 0 {
		limit = 50
	}

	// ユーザー検索を実行
	users, err := uc.userinfoRepository.SearchUsersByKeyword(keyword, limit)
	if err != nil {
		return nil, err
	}

	return users, nil
}
