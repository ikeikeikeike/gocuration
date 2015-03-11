package mailcatcher

import (
	"github.com/astaxie/beego/utils"
)

type CatchMailer struct {
	E *utils.Email
}

func NewCatchMailer() *CatchMailer {
	e := utils.NewEMail(`{
		"host":"127.0.0.1",
		"port":1025
	}`)
	e.From = "example@example.com"

	return &CatchMailer{E: e}
}

func (s *CatchMailer) SendMail(to, subject, htmlbody string) error {
	s.E.To = []string{to}
	s.E.Subject = subject
	s.E.HTML = htmlbody

	return s.E.Send()
}
