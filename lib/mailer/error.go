package mailer

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/astaxie/beego"
)

func SendStackTrace(msg []string) {
	for i := 1; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		msg = append(msg, fmt.Sprintf("%s:%d", file, line))
	}

	m := NewMailer()
	m.SendMail(
		beego.AppConfig.String("logMaildst"),
		fmt.Sprintf("[%s] Error mail", beego.RunMode),
		strings.Join(msg, "\n"),
	)
}
