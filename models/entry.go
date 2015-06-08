package models

import (
	"fmt"
	"strings"
	"time"

	"bitbucket.org/ikeikeikeike/antenna/lib/funcs"
	"bitbucket.org/ikeikeikeike/antenna/models/image"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/fatih/structs"
	"github.com/ikeikeikeike/gopkg/convert"
	"github.com/jinzhu/now"

	sani "github.com/kennygrant/sanitize"
	attr "github.com/oleiade/reflections"
)

type Entry struct {
	Id          int64     `orm:"auto" json:"id"`
	Url         string    `orm:"type(text);unique" json:"url"`
	Title       string    `orm:"size(255);null"    json:"title"`
	Content     string    `orm:"type(text);null"   json:"content"`
	SeoTitle    string    `orm:"size(255);null"    json:"-"`
	SeoContent  string    `orm:"type(text);null"   json:"-"`
	Encoded     string    `orm:"type(text);null"   json:"-"`
	Creator     string    `orm:"size(255);null"    json:"creator"`   // EntryCreator
	Publisher   string    `orm:"size(255);null"    json:"publisher"` // RSSFeeder
	PublishedAt time.Time `orm:"type(datetime);null;index" json:"publishedAt"`

	IsBan     string `orm:"default(none)" valid:"Required;Match(/^(none|soft|hard)$/)"`
	IsPenalty bool   `orm:"default(0)"`
	PageView  int64  `orm:"default(0);index"`

	Q string `orm:"type(text);null" json:"-"` // Gin Index

	Created time.Time `orm:"auto_now_add;type(datetime)" json:"created"`
	Updated time.Time `orm:"auto_now;type(datetime)" json:"updated"`

	Blog *Blog `orm:"rel(fk);index" json:"-"`

	Video   *Video   `orm:"reverse(one)" json:"-"`
	Picture *Picture `orm:"reverse(one)" json:"-"`
	Summary *Summary `orm:"reverse(one)" json:"-"`

	Tags   []*Tag   `orm:"rel(m2m);index;rel_through(bitbucket.org/ikeikeikeike/antenna/models.EntryTag)" json:"-"`
	Images []*Image `orm:"rel(m2m);index;rel_through(bitbucket.org/ikeikeikeike/antenna/models.EntryImage)" json:"-"`

	Scores []*Score `orm:"reverse(many)" json:"-"`
}

type EntryIndex struct {
	Entry

	Tags   []string `json:"tags"`
	Images []string `json:"images"`

	BlogName        string `json:"blogName"`
	BlogMediatype   string `json:"blogMediatype"`
	BlogAdsensetype string `json:"blogAdsensetype"`

	VideoDuration int      `json:"videoDuration"`
	VideoDivas    []string `json:"videoDivas"`
	VideoBracups  []string `json:"videoBracups"`
	VideoDomain   string   `json:"videoDomain"`

	PictureCharacters []string `json:"pictureCharacters"`
	PictureBracups    []string `json:"pictureBracups"`

	PictureAnime  string `json:"pictureAnime"`
	PictureAlias  string `json:"pictureAlias"`
	PictureAuthor string `json:"pictureAuthor"`
	PictureWorks  string `json:"pictureWorks"`

	// No real time score below.
	HatenaScore   int64 `json:"hatenaScore"`
	TwitterScore  int64 `json:"twitterScore"`
	FacebookScore int64 `json:"facebookScore"`

	// No real time score below.
	InScore  int64 `json:"inScore"`
	OutScore int64 `json:"outScore"`
}

func (m *Entry) IdStr() string {
	return convert.ToStr(m.Id)
}

// For elasticsearch index data
func (m *Entry) SearchData() *EntryIndex {
	tags := []string{}
	if m.Tags != nil {
		for _, tag := range m.Tags {
			tags = append(tags, tag.Name)
		}
	}

	images := []string{}
	if m.Images != nil && !funcs.IsImgFallback(m.Blog.Url) {
		for _, image := range m.Images {
			images = append(images, image.Src)
		}
	}
	if len(images) <= 0 {
		images = append(images, image.CachedRandomSrc("large"))
	}

	idx := &EntryIndex{
		Entry:  *m,
		Tags:   tags,
		Images: images,
	}

	if m.Blog != nil {
		idx.BlogName = m.Blog.Name
		idx.BlogMediatype = m.Blog.Mediatype
		idx.BlogAdsensetype = m.Blog.Adsensetype
	}

	if m.Video != nil {
		m.Video.LoadRelated()

		if m.Video.Divas != nil {
			for _, diva := range m.Video.Divas {
				idx.VideoDivas = append(idx.VideoDivas, diva.Name)
				idx.VideoBracups = append(idx.VideoBracups, diva.Bracup)
			}
		}

		if m.Video.Site != nil {
			idx.VideoDomain = m.Video.Site.Domain
		}

		idx.VideoDuration = m.Video.Duration
	}

	if m.Picture != nil {
		m.Picture.LoadRelated()

		for _, c := range m.Picture.Characters {
			idx.PictureCharacters = append(idx.PictureCharacters, c.Name)
			idx.PictureBracups = append(idx.PictureBracups, c.Bracup)
		}

		if m.Picture.Anime != nil {
			idx.PictureAnime = m.Picture.Anime.Name
			idx.PictureAlias = m.Picture.Anime.Alias
			idx.PictureAuthor = m.Picture.Anime.Author
			idx.PictureWorks = m.Picture.Anime.Works
		}
	}

	var (
		n string
		s *Score
	)

	if m.Scores != nil {
		for _, s = range m.Scores {
			n = fmt.Sprintf("%sScore", strings.Title(s.Name))
			attr.SetField(idx, n, s.Count)
		}
	}

	if m.Blog != nil {
		if m.Blog.Scores != nil {
			for _, s = range m.Blog.Scores {
				n = fmt.Sprintf("%sScore", strings.Title(s.Name))
				attr.SetField(idx, n, s.Count)
			}
		}
	}

	return idx
}

func (m *Entry) LoadRelated() *Entry {
	o := orm.NewOrm()
	_, _ = o.LoadRelated(m, "Blog")
	_, _ = o.LoadRelated(m, "Video")
	_, _ = o.LoadRelated(m, "Picture")
	// _, _ = o.LoadRelated(m, "Summary")
	_, _ = o.LoadRelated(m, "Tags", 2, DefaultPerEntities, 0, "-id")
	_, _ = o.LoadRelated(m, "Images", 2, DefaultPerEntities, 0, "-id")
	_, _ = o.LoadRelated(m, "Scores", 2, DefaultPerEntities)
	return m
}

func (m *Entry) RelLoader() {
	m.LoadRelated()
}

// For gin fulltext index: search source.
func (m *Entry) SetQ() *Entry {
	m.LoadRelated()

	arystr := structs.Values(m.SearchData())
	strs := make([]string, len(arystr))

	for i, s := range arystr {
		cl, _ := sani.HTMLAllowing(convert.ToStr(s))
		strs[i] = sani.HTML(cl)
	}

	m.Q = strings.Join(strs, " ")
	return m
}

func (m *Entry) TodayRanking(name string) *EntryRanking {
	cond := orm.NewCondition()
	n := now.New(time.Now().UTC())

	switch name {
	case "w", "weekly":
		cond = cond.And("begin_time", n.BeginningOfWeek()).And("begin_name", "weekly")
	case "m", "monthly":
		cond = cond.And("begin_time", n.BeginningOfMonth()).And("begin_name", "monthly")
	case "y", "yearly":
		cond = cond.And("begin_time", n.BeginningOfYear()).And("begin_name", "yearly")
	case "d", "dayly":
		fallthrough
	default:
		cond = cond.And("begin_time", n.BeginningOfDay()).And("begin_name", "dayly")
	}

	qs := orm.NewOrm().QueryTable("entry_ranking").
		SetCond(cond).
		Filter("entry", m.Id)

	var ranking EntryRanking
	err := qs.One(&ranking)
	if err != nil {
		return nil
	}

	return &ranking
}

// Is liveing entry?
func (m *Entry) IsLiving() bool {
	m.RelLoader()

	if m.Id <= 0 || m.Blog == nil {
		return false
	}

	return true
}

func (m *Entry) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Entry) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Entry) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(m, field, fields...)
}

func (m *Entry) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Entry) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func Entries() orm.QuerySeter {
	return orm.NewOrm().QueryTable("entry").OrderBy("-Id")
}

func PictureEntries() orm.QuerySeter {
	return orm.NewOrm().QueryTable("entry").Filter("picture__isnull", false).OrderBy("-Id")
}

func init() {
	orm.RegisterModelWithPrefix(
		beego.AppConfig.String("dbprefix"),
		new(Entry))
}

////////////// views method below.

func (m *Entry) CommaTags() string {
	tags := []string{}
	if m.Tags != nil {
		for _, tag := range m.Tags {
			tags = append(tags, tag.Name)
		}
	}
	return strings.Join(tags, ",")
}
