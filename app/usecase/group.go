package usecase

import (
	"backend/app/domain/entity"
	"backend/app/domain/repository"
	"errors"
)

type GroupUseCase struct {
	groupRepository repository.GroupRepository
}

func NewGroupUseCase(groupRepository repository.GroupRepository) *GroupUseCase {
	return &GroupUseCase{
		groupRepository: groupRepository,
	}
}

func (u *GroupUseCase) CreateGroup(name string, description string, leaderEmail string, icon string, groupSkills []string) (groupId string, err error) {
	group := entity.NewGroup(name, description, leaderEmail, icon)

	groupSkillsEntity := make([]entity.GroupSkill, len(groupSkills))
	for i, skill := range groupSkills {
		groupSkillsEntity[i] = *entity.NewGroupSkill(group.GroupID, skill)
	}

	err = u.groupRepository.InsertGroup(group, &groupSkillsEntity)
	if err != nil {
		return "", errors.New("failed to insert group")
	}
	return group.GroupID, nil
}
