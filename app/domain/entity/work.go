package entity

import (
	"backend/app/interfaces/request"

	"github.com/google/uuid"
)

type (
	WorkTable struct {
		ID          string `db:"id"`
		Title       string `db:"title"`
		Description string `db:"description"`
		URL         string `db:"url"`
		MovieUrl    string `db:"movie_url"`
		Security    int    `db:"security"`
	}
	Tag struct {
		ID     string `db:"id"`
		WorkID string `db:"work_id"`
		Tag    string `db:"tag"`
	}
	Image struct {
		ID     string `db:"id"`
		WorkID string `db:"work_id"`
		Image  string `db:"image_url"`
	}
	ReadWork struct {
		Title       string   `db:"title"`
		Description string   `db:"description"`
		ImageURLs   []string `db:"image_url"`
		Tags        []string `db:"tag"`
		WorkURL     string   `db:"url"`
		MovieUrl    string   `db:"movie_url"`
		Security    int      `db:"security"`
	}
	ReadWorksList struct {
		WorkID      string `db:"id"`
		Title       string `db:"title"`
		Images      string `db:"image_url"`
		Description string `db:"description"`
		Icon        string `db:"icon"`
	}
)

func NewWork(work request.CreateWorkRequest) (*WorkTable, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	body := WorkTable{
		ID:          u.String(),
		Title:       work.Title,
		Description: work.Description,
		URL:         work.WorkUrl,
		MovieUrl:    work.MovieUrl,
		Security:    work.Security,
	}
	return &body, nil
}

func NewWorkImages(work request.CreateWorkRequest, workID string) (*[]Image, error) {
	var images []Image
	for _, i := range work.Images {
		u, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}
		image := Image{
			ID:     u.String(),
			WorkID: workID,
			Image:  i.Image,
		}
		images = append(images, image)
	}
	return &images, nil
}

func NewWorkTags(work request.CreateWorkRequest, workID string) (*[]Tag, error) {
	var tags []Tag
	for _, i := range work.Tags {
		u, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}
		tag := Tag{
			ID:     u.String(),
			WorkID: workID,
			Tag:    i.Tag,
		}
		tags = append(tags, tag)
	}
	return &tags, nil
}

func NewReadWorksList(workID string, title string, images string, description string, icon string) *ReadWorksList {
	return &ReadWorksList{
		WorkID:      workID,
		Title:       title,
		Images:      images,
		Description: description,
		Icon:        icon,
	}
}
