package funcs

import "net/url"

func ToHost(u string) string {
	p, _ := url.Parse(u)
	return p.Host
}
