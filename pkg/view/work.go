package view

import "backend/pkg/model/dto"

type CreateWorkResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Images      []struct {
		Image string `json:"image"`
	} `json:"images"`
	URL  string `json:"URL"`
	Tags []struct {
		Tag string `json:"tag"`
	} `json:"tags"`
	Group    string `json:"group"`
	Security int    `json:"security"`
}

func ReturnCreateWork(work dto.CreateWorkRequest) CreateWorkResponse {
	return CreateWorkResponse{Title: work.Title, Description: work.Description, Images: work.Images, URL: work.URL, Tags: work.Tags, Group: work.Group, Security: work.Security}
}
