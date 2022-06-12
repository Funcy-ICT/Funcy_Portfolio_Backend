package dao

import (
	"backend/pkg/model/dto"
	"errors"
	"github.com/oklog/ulid"
	"log"
	"math/rand"
	"time"
)

const (
	//insert work
	InsertWorkQuery       = "INSERT INTO `works` (id,work_id,title,url,security) VALUES (?,?,?,?,?)"
	InsertWorkImagesQuery = "INSERT INTO `work_images` (id,image) VALUES (?,?)"
	InsertWorkTagsQuery   = "INSERT INTO `work_tags` (id,tag) VALUES (?,?)"
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

	log.Println(workID)

	stmt, err := Conn.Prepare(InsertWorkQuery)
	if err != nil {
		return "", err
	}
	_, err = stmt.Exec(userID, workID, workInfo.Title, workInfo.URL, workInfo.Security)
	if err != nil {
		log.Println(err)
		return "", errors.New("Unable to insert work data")
	}

	stmt, err = Conn.Prepare(InsertWorkImagesQuery)
	if err != nil {
		return "", err
	}
	for _, i := range workInfo.Images {
		_, err = stmt.Exec(workID, i.Image)
		if err != nil {
			return "", errors.New("Unable to insert work image data")
		}
	}

	stmt, err = Conn.Prepare(InsertWorkTagsQuery)
	if err != nil {
		return "", err
	}
	for _, i := range workInfo.Tags {
		_, err = stmt.Exec(workID, i.Tag)
		if err != nil {
			return "", errors.New("Unable to insert work tags data")
		}
	}

	return "", err
}
