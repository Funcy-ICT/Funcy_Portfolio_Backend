package repository

import "backend/app/domain/entity"

type GroupRepository interface {
	InsertGroup(group *entity.Group, groupSkills *[]entity.GroupSkill) error
}
