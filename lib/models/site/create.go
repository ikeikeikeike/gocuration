package site

import (
	"errors"

	"bitbucket.org/ikeikeikeike/antenna/models"
	"bitbucket.org/ikeikeikeike/antenna/models/site"
	"github.com/astaxie/beego"
)

/*
	Maybe creator for Video.Domain
*/
func ReadOrCreateByVideo(v *models.Video) (*models.Site, error) {
	var domains []string

	domains = site.SelectDomains(v.Url)
	if len(domains) <= 0 {
		domains = site.SelectDomains(v.Code)
		if len(domains) <= 0 {
			msg := "Could not set Domain to Video"
			return nil, errors.New(msg)
		}
	}

	// If domain nums than once
	if len(domains) > 1 {
		beego.Warning(
			"Maybe add domain(length=", len(domains),
			") to Video(Id=", v.Id, ")")
	}

	return site.ReadOrCreateByDomain(domains[0])
}
