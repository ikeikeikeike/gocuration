package site

import (
	"fmt"

	"bitbucket.org/ikeikeikeike/antenna/models"
	"github.com/astaxie/beego"
	"github.com/ikeikeikeike/favpath"
)

func Asg()           {}
func EroVideo()      {}
func Fc2()           {}
func JapanWhores()   {}
func Pornhost()      {}
func Pornhub()       {}
func Redtube()       {}
func TokyoTube()     {}
func TokyoPornTube() {}
func Tube8()         {}
func Xhamster()      {}
func Xvideos()       {}

func ReadOrCreateByDomain(domain string) (s *models.Site, err error) {
	s = &models.Site{Domain: domain}

	created, _, _ := s.ReadOrCreate("Domain")
	if created {
		finder := favpath.NewFinder()
		finder.Header("User-Agent", beego.AppConfig.String("UserAgent"))

		// TODO: SubDomainに対応されているか確認
		icon, err := finder.Find(fmt.Sprintf("http://%s", domain))
		if err != nil {
			msg := fmt.Sprintf(
				"finder.Find err Site{id=%d} icon=%s: ",
				s.Id, icon)
			beego.Warning(msg, err)
		}
		if icon != "" {
			if s.Icon == nil {
				s.Icon = &models.Image{}
				s.Icon.Insert()
			}
			s.Icon.Src = icon
			s.Icon.Update()
		}

		s.Update()
	}

	return
}
