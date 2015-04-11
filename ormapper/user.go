package ormapper

import "time"

type User struct {
	Id int64

	Email    string
	Password string

	Lastlogintime time.Time

	Created time.Time
	Updated time.Time

	Blogs []*Blog
}
