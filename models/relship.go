package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type EntryTag struct {
	Id int64 `orm:"auto"`

	Tag   *Tag   `orm:"rel(fk);index"`
	Entry *Entry `orm:"rel(fk);index"`
}

type EntryImage struct {
	Id int64 `orm:"auto"`

	Image *Image `orm:"rel(fk);index"`
	Entry *Entry `orm:"rel(fk);index"`
}

type VideoDiva struct {
	Id int64 `orm:"auto"`

	Diva  *Diva  `orm:"rel(fk);index"`
	Video *Video `orm:"rel(fk);index"`
}

type PictureCharacter struct {
	Id int64 `orm:"auto"`

	Character *Character `orm:"rel(fk);index"`
	Picture   *Picture   `orm:"rel(fk);index"`
}

func init() {
	orm.RegisterModelWithPrefix(
		beego.AppConfig.String("dbprefix"),
		new(EntryTag),
		new(EntryImage),
		new(VideoDiva),
		new(PictureCharacter),
	)
}
