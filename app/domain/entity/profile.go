package entity

type Profile struct {
	UserID          string `db:"user_id"`
	DisplayName     string `db:"display_name"`
	Icon            string `db:"icon"`
	HeaderImagePath string `db:"header_image"`
	Biography       string `db:"bio"`
}

func NewProfile(userID, displayName, icon, headerImagePath, bio string) *Profile {
	return &Profile{
		UserID:          userID,
		DisplayName:     displayName,
		Icon:            icon,
		HeaderImagePath: headerImagePath,
		Biography:       bio,
	}
}

func NewSkill(skillName, userID string) *Skill {
	return &Skill{
		SkillName: skillName,
		UserID:    userID,
	}
}

func NewSNS(userID, snsURL string) *SNS {
	return &SNS{
		UserID: userID,
		SnsURL: snsURL,
	}
}
