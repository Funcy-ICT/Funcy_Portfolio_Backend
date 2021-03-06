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
	InsertWorkQuery       = "INSERT INTO `works` (id,work_id,title,description,url,movie_url,security) VALUES (?,?,?,?,?,?,?)"
	InsertWorkImagesQuery = "INSERT INTO `work_images` (id,image_id,image_url) VALUES (?,?,?)"
	InsertWorkTagsQuery   = "INSERT INTO `work_tags` (id,tag_id,tag) VALUES (?,?,?)"
	//select work
	SelectWork = "SELECT works.title, works.description, works.url, works.security,works.movie_url, work_images.image_url, work_tags.tag from works inner join work_images on works.work_id = work_images.id inner join work_tags on works.work_id = work_tags.id where work_id = ?"
	//select works list
	SelectWorksList = "select works.work_id, works.title,image.image_url, users.icon from works inner join users on works.id = users.id inner join (SELECT image_url,id FROM work_images GROUP BY image_url,id)as image on works.work_id = image.id ORDER BY works.created_at DESC limit ?"
)

///post work
type createWork struct {
}

func MakeCreateWorkClient() createWork {
	return createWork{}
}

func (info *createWork) Request(userID string, workInfo dto.CreateWorkRequest) (string, error) {
	//ソート可能なuildを使用
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	workID := ulid.MustNew(ulid.Timestamp(t), entropy).String()
	log.Println(workInfo.Movie_url)
	stmt, err := Conn.Prepare(InsertWorkQuery)
	if err != nil {
		return "", err
	}
	_, err = stmt.Exec(userID, workID, workInfo.Title, workInfo.Description, workInfo.Movie_url, workInfo.Work_URL, workInfo.Security)
	if err != nil {
		log.Println(err)
		return "", errors.New("Unable to insert work data")
	}

	stmt, err = Conn.Prepare(InsertWorkImagesQuery)
	if err != nil {
		return "", err
	}
	for _, i := range workInfo.Images {
		imageID, err := uuid.NewRandom()
		if err != nil {
			return "", errors.New("userID generate is failed")
		}
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
		tagID, err := uuid.NewRandom()
		if err != nil {
			return "", errors.New("userID generate is failed")
		}
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
		log.Println(err)
		if err == sql.ErrNoRows {
			return rw, errors.New("not exist work data")
		}
	}
	defer rows.Close()
	for rows.Next() {
		work := &dto.WorkTable{}
		if err := rows.Scan(&work.Title, &work.Description, &work.URL, &work.Security, &work.Movie_url, &work.Image, &work.Tag); err != nil {
			return rw, errors.New("exist nil work data")
		}
		works = append(works, *work)
	}
	if len(works) == 0 {
		return rw, errors.New("work id is wrong")
	}

	err = rows.Err()
	if err != nil {
		log.Println(err)
		return rw, errors.New("err")
	}
	//response用に取得してきたデータを整形
	//image
	var sortImages []dto.Image
	image := dto.Image{Image: works[0].Image}
	//タグの格納
	checkTag := works[0].Tag
	var tags []dto.Tag
	tag := dto.Tag{Tag: works[0].Tag}
	tags = append(tags, tag)
	for i, w := range works {
		image = dto.Image{Image: w.Image}
		sortImages = append(sortImages, image)
		if i != 0 {
			if w.Tag != checkTag {
				tag = dto.Tag{Tag: w.Tag}
				tags = append(tags, tag)
				checkTag = w.Tag
			} else {
				continue
			}
		}
	}

	sort.Slice(tags, func(i, j int) bool {
		return tags[i].Tag < tags[j].Tag
	})

	sort.Slice(sortImages, func(i, j int) bool {
		return sortImages[i].Image < sortImages[j].Image
	})

	var images []dto.Image

	checkImage := sortImages[0]

	image = sortImages[0]
	images = append(images, image)
	for _, i := range sortImages {
		if i != checkImage {
			image = i
			images = append(images, image)
			checkImage = i
		} else {
			continue
		}
	}

	checkTags2 := tags[0]
	tags2 := tags[0]
	var sortTags []dto.Tag
	sortTags = append(sortTags, checkTags2)
	for _, i := range tags {
		log.Println(i)
		if i != checkTags2 {
			tags2 = i
			sortTags = append(sortTags, tags2)
			checkTags2 = i
		} else {
			continue
		}
	}

	w := dto.ReadWork{
		Title:       works[0].Title,
		Description: works[0].Description,
		URL:         works[0].URL,
		Images:      images,
		Movie_url:   works[0].Movie_url,
		Tags:        sortTags,
		Security:    works[0].Security,
	}
	return w, err
}

///get works list
type readWorksList struct {
}

func MakeReadWorksListClient() readWorksList {
	return readWorksList{}
}

var (
	rwl []dto.ReadWorksList
)

func (info *readWorksList) Request(number string) ([]dto.ReadWorksList, error) {

	var worksList []dto.ReadWorksList

	n := dto.S2i(number)

	rows, err := Conn.Query(SelectWorksList, n)
	if err == sql.ErrNoRows {
		return rwl, errors.New("not exist work data")
	}

	for rows.Next() {
		work := &dto.ReadWorksList{}
		if err := rows.Scan(&work.WorkID, &work.Title, &work.Images, &work.Icon); err != nil {
			log.Println(err)
			return rwl, errors.New("err")
		}
		worksList = append(worksList, *work)
	}

	return worksList, err
}
