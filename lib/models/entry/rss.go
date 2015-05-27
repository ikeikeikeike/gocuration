package entry

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"bitbucket.org/ikeikeikeike/antenna/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	rss "github.com/ikeikeikeike/go-pkg-rss"
	"github.com/ikeikeikeike/gopkg/convert"
	"github.com/ikeikeikeike/gopkg/extract/image"
	"github.com/ikeikeikeike/gopkg/str"

	sani "github.com/kennygrant/sanitize"
	// "github.com/k0kubun/pp"
)

var storedNum = 5

func AddsByChannels(chs []*rss.Channel, b *models.Blog) (entries []*models.Entry, errs []error) {
	o := orm.NewOrm()

	// DMCA affect
	num := (storedNum + b.VerifyScore())
	if b.IsPenalty {
		num = 1
	}
	if b.IsBan == "soft" {
		num = 1
	} else if b.IsBan == "hard" {
		num = 0
	}

	for _, ch := range chs {
		for idx, it := range ch.Items {
			if idx >= num {
				break
			}
			if exists := o.QueryTable("entry").Filter("Url", it.Links[0].Href).Exist(); !exists {
				s := new(models.Entry)
				if _, err := AddByItem(s, it, b); err != nil {
					errs = append(errs, err)
				} else {
					entries = append(entries, s)
				}
			}
		}
	}
	return
}

var reIgnoreSites = func() *regexp.Regexp {
	nameList := []string{
		`xvideos?`, `erovideo`, `redtube`,
		`hamster`, `fc2`, `videomega`,
		`download_button`, `erodougamo/10000`,
	}
	return regexp.MustCompile(fmt.Sprintf(`(?:%s)\.(?:jpe?g|png|gif)$`, strings.Join(nameList, `|`)))
}()

func AddByItem(s *models.Entry, it *rss.Item, b *models.Blog) (int64, error) {
	var (
		msg string
		tf  = "2006-01-02T15:04:05-07:00"
	)

	s.Blog = b
	s.Url = it.Links[0].Href
	s.Title = it.Title
	s.Creator = it.Creator
	s.Publisher = it.Publisher

	s.Encoded = it.Encoded
	s.Content = it.Description
	if s.Content == "" {
		s.Content = it.Encoded
	}
	s.Content, _ = sani.HTMLAllowing(s.Content)
	s.Content = sani.HTML(s.Content)

	s.SeoTitle = s.Title
	s.SeoContent = s.Content

	s.PublishedAt, _ = time.Parse(tf, it.Date)
	if it.PubDate != "" {
		s.PublishedAt, _ = time.Parse(tf, it.PubDate)
	}

	err := s.Insert() // Issue primary id

	// The following is relationship update.

	o := orm.NewOrm()

	var tags []*models.Tag
	for _, word := range convert.StrTo(it.Subject).MultiWord() {
		tag := &models.Tag{Name: word}
		_, _, _ = tag.ReadOrCreate("Name")
		tags = append(tags, tag)
	}
	_, err = o.QueryM2M(s, "Tags").Add(tags)
	if err != nil {
		beego.Warn("Error add tags:", err)
	}

	paths, err := image.ExtractImagePaths(it.Encoded)
	if err != nil {
		beego.Warn("ExtractImagePaths:", err)
	}

	if len(paths) > 0 {
		var imgs []*models.Image
		for _, src_ := range paths {
			src := str.Clean(src_)

			if reIgnoreSites.MatchString(src) {
				beego.Notice("ignore image: ", src)
				continue
			}

			info, err := image.NewFileInfo(src)
			if err != nil {
				beego.Warn("image.NewFileInfo:", err)
				continue
			} else if info.Width < 100 || info.Height < 100 {
				beego.Notice("image.NewFileInfo: less than 100px width/height")
				continue
			}

			img := &models.Image{
				Src:    src,
				Ext:    info.Ext,
				Mime:   info.Mime,
				Name:   info.Filename,
				Width:  info.Width,
				Height: info.Height,
			}
			_, _, _ = img.ReadOrCreate("Src")
			imgs = append(imgs, img)
		}

		if len(imgs) > 0 {
			num, err := o.QueryM2M(s, "Images").Add(imgs)
			if err != nil {
				msg = fmt.Sprintf(
					"%s Errors. Add images by Entry(id=%s):",
					s.Id, num)
				beego.Warn(msg, err)
				return s.Id, err
			}
		}
	}

	UpdateSocials(s)

	return s.Id, s.SetQ().Update()
}
