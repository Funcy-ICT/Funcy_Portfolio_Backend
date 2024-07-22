package request

type CreateGroupRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	LeaderEmail string   `json:"leader_email"`
	Icon        string   `json:"icon"`
	GroupSkills []string `json:"group_skills"`
}
