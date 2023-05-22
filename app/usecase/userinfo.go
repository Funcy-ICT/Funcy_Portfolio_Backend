package usecase

import (
	"backend/app/domain/entity"
	"backend/app/domain/repository"
)

type UserinfoUseCase struct {
	profileRepository repository.ProfileRepository
	groupRepository   repository.GroupRepository
	skillRepository   repository.SkillRepository
	snsRepository     repository.SNSRepository
	workRepository    repository.WorkRepository
}

func NewUserinfoUsecace(
	profileRepository repository.ProfileRepository,
	groupRepository repository.GroupRepository,
	skillRepository repository.SkillRepository,
	snsRepository repository.SNSRepository,
	workRepository repository.WorkRepository,
) *UserinfoUseCase {
	return &UserinfoUseCase{
		profileRepository: profileRepository,
		groupRepository:   groupRepository,
		skillRepository:   skillRepository,
		snsRepository:     snsRepository,
		workRepository:    workRepository,
	}
}

func (u *UserinfoUseCase) GetUserinfo(userID string) (*entity.Userinfo, error) {
	userProfile, err := u.profileRepository.GetProfile(userID)
	if err != nil {
		return nil, err
	}

	joinedGroups, err := u.groupRepository.GetJoinedGroups(userID)
	if err != nil {
		return nil, err
	}

	skills, err := u.skillRepository.GetSkills(userID)
	if err != nil {
		return nil, err
	}

	sns, err := u.snsRepository.GetSNS(userID)
	if err != nil {
		return nil, err
	}

	works, err := u.workRepository.SelectWorksByUserID(userID)
	if err != nil {
		return nil, err
	}

	return &entity.Userinfo{
		Profile: userProfile,
		Groups:  joinedGroups,
		Skills:  skills,
		SNS:     sns,
		Works:   works,
	}, nil
}
