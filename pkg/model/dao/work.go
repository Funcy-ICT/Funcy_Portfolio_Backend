package dao

import (
	"backend/pkg/model/dto"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/oklog/ulid"
	"log"
	"math/rand"
	"sort"
	"time"
)

const (
	//insert work
	InsertWorkQuery       = "INSERT INTO `works` (id,work_id,title,description,url,security) VALUES (?,?,?,?,?,?)"
	InsertWorkImagesQuery = "INSERT INTO `work_images` (id,image_id,image_url) VALUES (?,?,?)"
	InsertWorkTagsQuery   = "INSERT INTO `work_tags` (id,tag_id,tag) VALUES (?,?,?)"
	//select work
	SelectWork = "SELECT works.title, works.description, works.url, works.security, work_images.image_url, work_tags.tag from works inner join work_images on works.id = work_images.id inner join work_tags on work_images.id = work_tags.id where works.id = ?"
)

///post work
type createWork struct {
}

func MakeCreateWorkClient() createWork {
	return createWork{}
}

func (info *createWork) Request(userID string, workInfo dto.CreateWorkRequest) (string, error) {

	imageID, err := uuid.NewRandom()
	tagID, err := uuid.NewRandom()

	log.Println(userID)
	if err != nil {
		return "", errors.New("userID generate is failed")
	}
	log.Println(userID)
	if err != nil {
		return "", errors.New("userID generate is failed")
	}
	//ソート可能なuildを使用
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	workID := ulid.MustNew(ulid.Timestamp(t), entropy).String()

	log.Println(workID)

	stmt, err := Conn.Prepare(InsertWorkQuery)
	if err != nil {
		return "", err
	}
	_, err = stmt.Exec(userID, workID, workInfo.Title, workInfo.Description, workInfo.URL, workInfo.Security)
	if err != nil {
		log.Println(err)
		return "", errors.New("Unable to insert work data")
	}

	stmt, err = Conn.Prepare(InsertWorkImagesQuery)
	if err != nil {
		return "", err
	}
	for _, i := range workInfo.Images {
		_, err = stmt.Exec(workID, imageID, i.Image)
		if err != nil {
			return "", errors.New("Unable to insert work image data")
		}
	}

	stmt, err = Conn.Prepare(InsertWorkTagsQuery)
	if err != nil {
		return "", err
	}
	for _, i := range workInfo.Tags {
		_, err = stmt.Exec(workID, tagID, i.Tag)
		if err != nil {
			return "", errors.New("Unable to insert work tags data")
		}
	}

	return "", err
}

///get work
type readWork struct {
}

func MakeReadWorkClient() readWork {
	return readWork{}
}

var (
	rw  dto.ReadWork
	cwr dto.CreateWorkRequest
)

func (info *readWork) Request(workID string) (dto.ReadWork, error) {

	var works []dto.WorkTable

	rows, err := Conn.Query(SelectWork, workID)
	if err != nil {
		if err == sql.ErrNoRows {
			return rw, errors.New("not exist work data")
		}
	}

	//取得してきた複数(単数)のレコード1つずつ処理
	for rows.Next() {
		var work dto.WorkTable
		if err := rows.Scan(&work.Title, &work.Description, &work.URL, &work.Security, &work.Image, &work.Tag); err != nil {

			if err == sql.ErrNoRows {
				return rw, errors.New("err")
			}
		}
		works = append(works, work)
	}

	//response用に取得してきたデータを整形
	//image
	var sortImages []string
	//タグの格納
	checkTag := works[0].Tag
	var tags []dto.Tag
	tag := dto.Tag{Tag: works[0].Tag}
	tags = append(tags, tag)
	for i, w := range works {
		if i != 0 {
			if w.Tag != checkTag {
				tag = dto.Tag{Tag: w.Tag}
				tags = append(tags, tag)
			}
		}
		checkTag = w.Tag
		sortImages = append(sortImages, w.Image)
	}
	sort.Slice(sortImages, func(i, j int) bool {
		return sortImages[i] < sortImages[j]
	})

	checkImage := sortImages[0]
	var images []dto.Image
	image := dto.Image{Image: works[0].Image}
	images = append(images, image)
	for _, i := range sortImages {
		if i != checkImage {
			image = dto.Image{Image: i}
			images = append(images, image)
		}
		checkImage = i
	}

	w := dto.ReadWork{
		Title:       works[0].Title,
		Description: works[0].Description,
		URL:         works[0].URL,
		Images:      images,
		Tags:        tags,
		Security:    works[0].Security,
	}

	return w, err
}
