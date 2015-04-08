package ormapper

import (
	"database/sql"
	"time"
)

type Summary struct {
	Id int64

	Sort int64

	Created time.Time
	Updated time.Time

	Entry   Entry
	EntryId sql.NullInt64

	Scores []Score
}
