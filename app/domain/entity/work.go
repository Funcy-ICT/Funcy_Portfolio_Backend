package entity

import "github.com/google/uuid"

type (
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
		Thumbnail   string   `db:"thumbnail"`
		UserId      string   `db:"user_id"`
		ImageURLs   []string `db:"image_url"`
		Tags        []string `db:"tag"`
		WorkUrl     string   `db:"url"`
		MovieUrl    string   `db:"movie_url"`
		GroupID     string   `db:"group_id"`
		Security    int      `db:"security"`
	}
	ReadWorksList struct {
		WorkID      string `db:"id"`
		Title       string `db:"title"`
		Thumbnail   string `db:"thumbnail"`
		Description string `db:"description"`
		Icon        string `db:"icon"`
		Security    int    `db:"security"`
	}
	InsertWork struct {
		ID          string `db:"id"`
		Title       string `db:"title"`
		Description string `db:"description"`
		Thumbnail   string `db:"thumbnail"`
		WorkUrl     string `db:"work_url"`
		MovieUrl    string `db:"movie_url"`
		GroupID     string `db:"group_id"`
		Security    int    `db:"security"`
	}
	UpdateWork struct {
		ID          string `db:"id"`
		Title       string `db:"title"`
		Description string `db:"description"`
		Thumbnail   string `db:"thumbnail"`
		WorkUrl     string `db:"work_url"`
		MovieUrl    string `db:"movie_url"`
		GroupID     string `db:"group_id"`
		Security    int    `db:"security"`
	}
)

func NewInsertWork(
	title string,
	description string,
	thumbnail string,
	workUrl string,
	movieUrl string,
	groupID string,
	security int) InsertWork {
	return InsertWork{
		ID:          uuid.NewString(),
		Title:       title,
		Description: description,
		Thumbnail:   thumbnail,
		WorkUrl:     workUrl,
		MovieUrl:    movieUrl,
		GroupID:     groupID,
		Security:    security,
	}
}

func NewUpdateWork(
	id string,
	title string,
	description string,
	thumbnail string,
	workUrl string,
	movieUrl string,
	groupID string,
	security int) UpdateWork {
	return UpdateWork{
		ID:          id,
		Title:       title,
		Description: description,
		Thumbnail:   thumbnail,
		WorkUrl:     workUrl,
		MovieUrl:    movieUrl,
		GroupID:     groupID,
		Security:    security,
	}
}

func NewImage(workID string, image string) Image {
	return Image{
		ID:     uuid.NewString(),
		WorkID: workID,
		Image:  image,
	}
}

func NewTag(workID string, tag string) Tag {
	return Tag{
		ID:     uuid.NewString(),
		WorkID: workID,
		Tag:    tag,
	}
}
