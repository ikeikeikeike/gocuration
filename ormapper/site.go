package ormapper

import (
	"database/sql"
	"time"
)

type Site struct {
	Id int64

	Name   string
	Domain string

	Outline string

	Created time.Time
	Updated time.Time

	Icon   Image
	IconId sql.NullInt64

	Videos []Video
}
