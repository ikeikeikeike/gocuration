package ranking

import (
	"time"

	"bitbucket.org/ikeikeikeike/antenna/models"
)

func ReadOrCreateEntry(id int64, n string, t time.Time) (bool, *models.EntryRanking, error) {
	e := &models.Entry{Id: id}
	e.Read()

	c := &models.EntryRanking{BeginName: n, BeginTime: t, Entry: e}
	created, _, err := c.ReadOrCreate("BeginName", "BeginTime", "Entry")
	return created, c, err
}

func ReadOrCreateVideo(id int64, n string, t time.Time) (bool, *models.VideoRanking, error) {
	e := &models.Entry{Id: id}
	e.Read()
	e.LoadRelated()

	c := &models.VideoRanking{BeginName: n, BeginTime: t, Video: e.Video}
	created, _, err := c.ReadOrCreate("BeginName", "BeginTime", "Video")
	return created, c, err
}

func ReadOrCreatePicture(id int64, n string, t time.Time) (bool, *models.PictureRanking, error) {
	e := &models.Entry{Id: id}
	e.Read()
	e.LoadRelated()

	c := &models.PictureRanking{BeginName: n, BeginTime: t, Picture: e.Picture}
	created, _, err := c.ReadOrCreate("BeginName", "BeginTime", "Picture")
	return created, c, err
}
