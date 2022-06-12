package dto

import (
	"database/sql"
	"log"
)

//作品投稿リクエスト

type CreateWorkRequest struct {
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

type WorkTable struct {
	Title       string
	Description string
	Image       string
	URL         string
	Tag         string
	Security    int
}

type ReadWork struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Images      []Image `json:"images"`
	URL         string  `json:"URL"`
	Tags        []Tag   `json:"tags"`
	Security    int     `json:"security"`
}

type Tag struct {
	Tag string
}
type Image struct {
	Image string
}

func ConvertToWork(row *sql.Rows) (*WorkTable, error) {
	var work WorkTable
	if err := row.Scan(&work.Title, &work.Description, &work.Image, &work.URL, &work.Tag, work.Security); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		log.Println(err)
		return nil, err
	}
	return &work, nil
}
