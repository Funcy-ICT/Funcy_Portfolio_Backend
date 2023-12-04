package usecase

import (
	"backend/app/domain/entity"
	"backend/app/domain/repository"
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
	groupSkillsEntity := []entity.GroupSkill{}
	for _, skill := range groupSkills {
		groupSkillsEntity = append(groupSkillsEntity, *entity.NewGroupSkill(group.GroupID, skill))
	}
	err = u.groupRepository.InsertGroup(group, &groupSkillsEntity)
	if err != nil {
		return "", err
	}
	return group.GroupID, nil
}
