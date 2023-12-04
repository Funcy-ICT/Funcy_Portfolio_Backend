package infrastructure

import (
	"backend/app/domain/entity"
	"backend/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type GroupRepositoryImpl struct {
	db *sqlx.DB
}

func NewGroupRepository(db *sqlx.DB) repository.GroupRepository {
	return &GroupRepositoryImpl{db: db}
}

func (gr *GroupRepositoryImpl) InsertGroup(group *entity.Group, groupSkills *[]entity.GroupSkill) error {
	tx, err := gr.db.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO groups (id, name, description, leader_email, icon) VALUES (?,?,?,?,?)",
		group.GroupID, group.Name, group.Description, group.LeaderEmail, group.Icon)
	if err != nil {
		return err
	}

	for _, groupSkill := range *groupSkills {
		_, err := tx.Exec("INSERT INTO group_skills (group_id, skill_name) VALUES (?,?)", groupSkill.GroupID, groupSkill.SkillName)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
