package accessctl

import (
	"fmt"
	neturl "net/url"
	"regexp"
	"strings"
	"time"

	"github.com/astaxie/beego"
	ctl "github.com/ikeikeikeike/gopkg/accessctl"
)

/*
	以下のいずれかを満たす場合はAccessLogとして集計しない

		- 正しいUserAgentか
		- Refererが存在するか
		- 自サイトのRefererではないか(conf.httphostを使用)
		- 自サイトのRefererではないか(RefererとrequestURLのドメインが同一)
*/
func IsInAccessLog(in *InLog) bool {

	if !ctl.IsUA(in.UserAgent) {
		beego.Debug("Valid UserAgent")
		return false
	} else if in.Referer == "" {
		beego.Debug("blank referer")
		return false
	} else if strings.Contains(in.Referer, beego.AppConfig.String("domain")) {
		beego.Debug("domain in referer", in.Referer, beego.AppConfig.String("domain"))
		return false
	}

	ru, err := neturl.Parse(in.Referer)
	if err != nil {
		return true
	}
	uu, err := neturl.Parse(in.Url)
	if err != nil {
		return true
	}

	if ru.Host == uu.Host {
		beego.Debug("same domain")
		return false
	}

	return true
}

type ScoringInLog struct {
	InLogs  []*InLog
	Cleaned bool
}

func NewScoringInLog() *ScoringInLog {
	return &ScoringInLog{Cleaned: true}
}

func (sin *ScoringInLog) Bootstrap() (err error) {

	counter, err := ctl.NewCounter()
	if err != nil {
		return
	}

	logs := counter.Listall()

	if sin.Cleaned {
		counter.Clean()
	}

	for _, l := range logs {
		sin.InLogs = append(sin.InLogs, BytesToInLog([]uint8(l)))
	}

	return
}

func (sin *ScoringInLog) Find(key string) (inlogs []*InLog, err error) {
	re := regexp.MustCompile(key)

	for _, inlog := range sin.InLogs {
		if !IsInAccessLog(inlog) {
			continue
		} else if len(re.FindAllString(inlog.Referer, -1)) < 1 {
			continue
		}

		inlogs = append(inlogs, inlog)
	}

	return
}

/*
	- IsInAccessLogで非カウント対象を弾いている
	- UserAgent, RemoteAddr, 1時間で UU(1セッション)とする
*/
func (sin *ScoringInLog) Scoring(url string) (score int, err error) {
	var key string = url

	u, err := neturl.Parse(url)
	if err == nil {
		domain := strings.Split(u.Host, ":")[0]
		if strings.Contains(domain, "livedoor.jp") {
			key = strings.Split(u.Path, "/")[1]
		} else {
			key = domain
		}
	}

	inlogs, err := sin.Find(key)
	if err != nil {
		return
	}

	scores := make(map[string]int)

	for _, il := range inlogs {
		t, _ := time.Parse(Timeformat, il.Time)

		scores[fmt.Sprintf("%s||%s||%s",
			il.RemoteHost, il.UserAgent, t.Format("2006-01-02T15"))]++
	}

	return len(scores), err
}
