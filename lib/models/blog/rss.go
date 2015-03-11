package blog

import (
	"fmt"

	"bitbucket.org/ikeikeikeike/antenna/models"

	"github.com/astaxie/beego"

	"github.com/ikeikeikeike/favpath"
	rss "github.com/ikeikeikeike/go-pkg-rss"
	"github.com/ikeikeikeike/gopkg/str"

	// "github.com/k0kubun/pp"
)

func UpdateByChannel(b *models.Blog, ch *rss.Channel, force bool) error {
	var (
		msg  string
		href string = str.Clean(ch.Links[0].Href)
	)

	if b.Url == "" || b.Url != href || force {
		b.Url = href
	}
	if b.Name == "" || force {
		b.Name = ch.Title
	}

	if b.Icon == nil {
		b.Icon = &models.Image{}
		b.Icon.Insert()
	}

	if b.Icon.Src == "" || force {
		finder := favpath.NewFinder()
		finder.Header("User-Agent", beego.AppConfig.String("UserAgent"))

		// c := s3.NewNamedClient().
		// SetAuth(beego.AppConfig.String("AwsAccessKey"), beego.AppConfig.String("AwsSecretKey")).
		// SetBucket(beego.AppConfig.String("AwsS3Bucket"))

		icon, err := finder.Find(ch.Links[0].Href)
		if err != nil {
			msg = fmt.Sprintf(
				"finder.Find err Blog{id=%d} icon=%s: ", b.Id, icon)
			beego.Warn(msg, err)
			// } else if icon, err = c.PutByUrl(icon); err != nil {
			// beego.Warn("s3 upload err Blog{id=%s} icon=%s: %s", b.Id, icon, err)
		}
		if icon != "" {
			b.Icon.Src = icon
			b.Icon.Update()
		}
	}
	// if b.LastModified != "" || force { b.LastModified = time.Now() }
	return b.Update()
}
