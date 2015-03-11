package accessctl

import "encoding/json"

var Timeformat = "2006-01-02T15:04:05Z07:00"

type InLog struct {
	RemoteHost string `json: "remoteHost"`
	Time       string `json: "time"`
	Method     string `json: "method"`
	Url        string `json: "url"`
	Status     string `json: "status"`
	Referer    string `json: "referer"`
	UserAgent  string `json: "userAgent"`
}

func BytesToInLog(bytes []byte) (inlog *InLog) {
	json.Unmarshal(bytes, &inlog)
	return
}
