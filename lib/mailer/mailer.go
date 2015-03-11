package mailer

import (
	"github.com/astaxie/beego"
	ses "github.com/ikeikeikeike/gopkg/mailer"
	catch "bitbucket.org/ikeikeikeike/antenna/lib/mailer/mailcatcher"
)

type Mailer interface {
	SendMail(to, subject, htmlbody string) error
}

func NewMailer() Mailer {
	if beego.RunMode == "prod" {
		return ses.NewSesMailer()
	} else {
		return catch.NewCatchMailer()
	}
}
