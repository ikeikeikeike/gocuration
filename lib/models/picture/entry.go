package picture

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	libm "bitbucket.org/ikeikeikeike/antenna/lib/models"
	libentry "bitbucket.org/ikeikeikeike/antenna/lib/models/entry"
	"bitbucket.org/ikeikeikeike/antenna/models"
	"bitbucket.org/ikeikeikeike/antenna/models/anime"
	"bitbucket.org/ikeikeikeike/antenna/models/character"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/ikeikeikeike/gopkg/extract/image"
	"github.com/ikeikeikeike/gopkg/str"
)

var (
	ReWords *regexp.Regexp
	ReAlias *regexp.Regexp

	// 曖昧な言葉
	IgnoreWords = []string{
		"小悪魔", "女王様",
		"主人公", "姉さん",
		"篠崎愛", "生徒たち",
		"校長先生", "女の子",
		"女教師", "女子高生",
		"女子大生", "小学生",
		"幼稚園児", "看護婦",
		"TEST", "test",
		"女学生", "お嬢さん",
		"未亡人", "管理人",
		"委員長", "学級長",
		"XXX", "MIX",
		"new", "男子生徒",
		"佐々木", "Kir",
		"お婆ちゃん", "小早川",
		"eve", "EVE",
		"ママさん", "AKI",
		"Spy", "王子様",
		"救世主", "佐久間",
		"扇風機", "暗殺者",
		"学園長", "殺人犯",
	}

	// 曖昧な言葉
	IgnoreAlias = []string{
		"パンスト", "ガールフレンド（仮）",
		"めぐみ", "タッチ", "みゆき", "かりん",
	}
)

func init() {
	ReWords = regexp.MustCompile(fmt.Sprintf(`^(%s)$`, strings.Join(IgnoreWords, "|")))
	ReAlias = regexp.MustCompile(fmt.Sprintf(`^(%s)$`, strings.Join(IgnoreAlias, "|")))
}

/*
	Added models below ::

		- Picture
		- []Picture.Image
		- Site
		- Caracter
		- Anime

	.. note:: only create

*/
func AddsByEntries(entries []*models.Entry) (errs []error) {
	var (
		err error
		p   *models.Picture
	)

	o := orm.NewOrm()

	info := image.NewInfo()
	info.Header("User-Agent", beego.AppConfig.String("UserAgent"))

	for _, e := range entries {
		e.RelLoader()

		ext := libentry.NewExtractor()
		ext.Header("User-Agent", beego.AppConfig.String("UserAgent"))

		if err := ext.Do(e); err != nil {
			beego.Warning("Error Entry{id=", e.Id, "}:", err)
			errs = append(errs, err)
			continue
		}

		// If has 500 entities, We presume called from user command.
		if len(entries) > 500 {
			time.Sleep(41 * time.Second)
		}

		p = e.Picture
		if p == nil {
			// Save Picture: last once value in slice...
			p = new(models.Picture)
			p.Entry = e

			// Need video's primary id
			p.Insert()

			for _, img := range ext.Imgs() {
				if beego.AppConfig.String("runmode") == "prod" {
					time.Sleep(300 * time.Millisecond)
				}

				src := str.Clean(img.Src)

				f, err := info.Fetch(src)
				if err != nil {
					beego.Notice("image.NewFileInfo:", err)
					continue
				} else if f.Width < 400 || f.Height < 330 {
					beego.Notice("fileinfo: less than 400/330(width/height)")
					continue
				}

				img := &models.Image{
					Src: src, Name: img.Alt, Picture: p, Ext: f.Ext,
					Mime: f.Mime, Width: f.Width, Height: f.Height,
				}

				img.Insert()
			}
		}
		p.RelLoader()

		// Save Character
		var its []*models.Character
		for _, c := range character.CachedCharacters() {
			if libm.ReHK3.MatchString(c.Name) {
				continue
			} else if ReWords.MatchString(c.Name) {
				continue
			} else if c.Id > 0 && len([]rune(c.Name)) > 2 {
				if strings.Contains(e.Title, c.Name) {
					its = append(its, c)
				} else if strings.Contains(e.Content, c.Name) {
					its = append(its, c)
				} else {
					for _, t := range e.Tags {
						if strings.Contains(t.Name, c.Name) {
							its = append(its, c)
						}
					}
				}
			}
		}
		for _, it := range its {
			if !o.QueryM2M(p, "Characters").Exist(it) {
				_, err = o.QueryM2M(p, "Characters").Add(it)
				if err != nil {
					beego.Warn("Warn add characters:", err)
				}
			}
		}

		// Save Anime
		for _, c := range anime.CachedAnimes() {

			if c.Id > 0 {
				for _, aka := range strings.Split(c.Alias, ",") {
					if len([]rune(aka)) > 2 && !ReAlias.MatchString(aka) {
						if strings.Contains(e.Title, aka) {
							p.Anime = c
						} else if strings.Contains(e.Content, aka) {
							p.Anime = c
						}
					}
				}

				if len([]rune(c.Name)) > 2 {
					if strings.Contains(e.Title, c.Name) {
						p.Anime = c
					} else if strings.Contains(e.Content, c.Name) {
						p.Anime = c
					} else {
						for _, t := range e.Tags {
							if strings.Contains(t.Name, c.Name) {
								p.Anime = c
							}
						}
					}
				}

				if p.Anime != nil {
					p.Update("Anime")
				}
			}
		}

	}
	return
}
