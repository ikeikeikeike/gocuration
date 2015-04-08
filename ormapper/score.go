package ormapper

import (
	"database/sql"
	"time"
)

type Score struct {
	Id    int64
	Name  string
	Count int64

	Created time.Time
	Updated time.Time

	Blog   Blog
	BlogId sql.NullInt64

	Entry   Entry
	EntryId sql.NullInt64

	Summary   Summary
	SummaryId sql.NullInt64
}
