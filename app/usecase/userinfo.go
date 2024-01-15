package usecase

import (
	"backend/app/domain/entity"
	"backend/app/domain/repository"
	"backend/app/interfaces/request"
)

type UserinfoUseCase struct {
	userinfoRepository repository.UserinfoRepository
	workRepository     repository.WorkRepository
}

func NewUserinfoUsecace(
	userinfoRepository repository.UserinfoRepository,
	workRepository repository.WorkRepository,
) *UserinfoUseCase {
	return &UserinfoUseCase{
		userinfoRepository: userinfoRepository,
		workRepository:     workRepository,
	}
}

func (u *UserinfoUseCase) GetUserinfo(userID string) (*entity.Userinfo, *[]*entity.ReadWorksList, error) {
	userinfo, err := u.userinfoRepository.SelectUserinfoByUserID(userID)
	if err != nil {
		return nil, nil, err
	}

	works, err := u.workRepository.SelectWorksByUserID(userID)
	if err != nil {
		return nil, nil, err
	}

	return userinfo, works, err
}

func (u *UserinfoUseCase) UpdateUserinfo(userID string, userinfo *request.UpdateUserInfo) error {
	parsed := new(entity.UpdateUserinfo)
	parsed.Profile = &entity.Profile{
		UserID:          userID,
		DisplayName:     userinfo.DisplayName,
		Icon:            userinfo.Icon,
		HeaderImagePath: userinfo.HeaderImagePath,
		Biography:       userinfo.Bio,
	}
	parsed.Skills = &[]*entity.Skill{}
	for _, s := range userinfo.Skills {
		*parsed.Skills = append(*parsed.Skills, &entity.Skill{SkillName: s, UserID: userID})
	}
	parsed.SNS = &[]*entity.SNS{}
	for _, s := range userinfo.SNS {
		*parsed.SNS = append(*parsed.SNS, &entity.SNS{SnsURL: s, UserID: userID})
	}

	return u.userinfoRepository.UpdateUserinfo(parsed)
}

func (u *UserinfoUseCase) CreateUserinfo(userID string, userinfo *request.UserInfo) error {
	parsed := new(entity.Userinfo)
	parsed.Profile = &entity.Profile{
		UserID:          userID,
		DisplayName:     userinfo.DisplayName,
		Icon:            userinfo.Icon,
		HeaderImagePath: userinfo.HeaderImagePath,
		Biography:       userinfo.Bio,
	}
	parsed.JoinedGroups = &[]*entity.GroupMember{}
	parsed.Skills = &[]*entity.Skill{}
	parsed.SNS = &[]*entity.SNS{}

	return u.userinfoRepository.CreateNewUserinfo(parsed)
}
