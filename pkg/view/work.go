package view

import "backend/pkg/model/dto"

type CreateWorkResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Images      []struct {
		Image string `json:"image"`
	} `json:"images"`
	URL       string `json:"URL"`
	Movie_url string `json:"movie_url"`
	Tags      []struct {
		Tag string `json:"tag"`
	} `json:"tags"`
	Group    string `json:"group"`
	Security int    `json:"security"`
}

func ReturnCreateWork(work dto.CreateWorkRequest) CreateWorkResponse {
	return CreateWorkResponse{Title: work.Title, Description: work.Description, Images: work.Images, URL: work.Work_URL, Movie_url: work.Movie_url, Tags: work.Tags, Group: work.Group, Security: work.Security}
}

type ReadWorkResponse struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Images      []Image `json:"images"`
	URL         string  `json:"URL"`
	Tags        []Tag   `json:"tags"`
	Security    int     `json:"security"`
}

type Tag struct {
	Tag string `json:"tag"`
}
type Image struct {
	Image string `json:"image"`
}

//func ReturnReadWork(work dto.ReadWork) ReadWorkResponse {
//	return ReadWorkResponse{Title: work.Title, Description: work.Description, Images: work.Images, URL: work.URL, Tags: work.Tags, Security: work.Security}
//}
