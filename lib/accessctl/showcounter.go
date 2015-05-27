package accessctl

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/ikeikeikeike/gopkg/convert"
	"github.com/ikeikeikeike/gopkg/redis"
)

type ShowCounter struct {
	rc      *redis.Client
	baseKey string

	Cleaned bool
}

func NewShowCounter() *ShowCounter {
	return &ShowCounter{
		rc:      redis.NewClient(),
		baseKey: "accessctl.showcounter:",
		Cleaned: true,
	}
}

func (c *ShowCounter) Bootstrap() (err error) {
	config := fmt.Sprintf(`{"conn": "%s"}`, beego.AppConfig.String("RedisConn"))
	err = c.rc.Initialize(config)
	return
}

func (c *ShowCounter) Clean(path string) {
	c.rc.Del(path)
}

func (c *ShowCounter) Findlist(path string) (list []string) {
	s := reflect.ValueOf(c.rc.Listall(path))

	list = make([]string, s.Len())
	for i := 0; i < s.Len(); i++ {
		list[i] = convert.ToStr(s.Index(i).Elem().Bytes())
	}
	return
}

type ShowResult struct {
	ID    int64
	Count int64
}

func (c *ShowCounter) Counting(key string) (r []ShowResult, err error) {
	k := c.baseKey + "/" + key + "/"

	var keys interface{}
	if keys, err = c.rc.Do("KEYS", k+"*"); err != nil {
		return
	}

	s := reflect.ValueOf(keys)
	names := make([]string, s.Len())
	for i := 0; i < s.Len(); i++ {
		names[i] = convert.ToStr(s.Index(i).Elem().Bytes())
	}

	for _, name := range names {
		logs := c.Findlist(name)

		if c.Cleaned {
			c.Clean(name)
		}

		var inLogs []*InLog
		for _, log := range logs {
			inLogs = append(inLogs, BytesToInLog([]uint8(log)))
		}

		scores := make(map[string]int)
		for _, il := range inLogs {
			t, _ := time.Parse(time.RFC1123, il.Time)

			scores[fmt.Sprintf("%s||%s||%s",
				il.RemoteHost, il.UserAgent, t.Format("2006-01-02T15:04"))]++
		}

		id, err := convert.StrTo(strings.Replace(name, k+"v", "", -1)).Int64()
		if err != nil {
			continue
		}
		r = append(r, ShowResult{ID: id, Count: int64(len(scores))})
	}

	return
}
