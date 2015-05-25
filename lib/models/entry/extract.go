package entry

import (
	"bytes"
	"fmt"
	urlparse "net/url"
	"regexp"
	"strings"

	libm "bitbucket.org/ikeikeikeike/antenna/lib/models"
	gq "github.com/PuerkitoBio/goquery"
	behavior "github.com/ikeikeikeike/gopkg/net/http"
	"github.com/ikeikeikeike/gopkg/str"

	"golang.org/x/net/html"

	"bitbucket.org/ikeikeikeike/antenna/models"
)

var (
	defaultUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_5) AppleWebKit/600.3.18 (KHTML, like Gecko) Version/7.1.3 Safari/537.85.12"

	embedUrls = []string{
		`flashservice.xvideos.com/embedframe`,
		`video.fc2.com/content`,
		`video.fc2.com/a/content`,
		`video.fc2.com/ja/a/content`,
		`www.tokyo-tube.com/embedcode`,
		// ero-video
		`xhamster.com/xembed.php`,
		// asg
		`www.pornhost.com/embed`,
		`www.tube8.com/embed/asian`,
		// redtube
		`jp.pornhub.com/embed`,
		`www.pornhub.com/embed`,
		`japan-whores.com/embed`,
		`www.tokyo-porn-tube.com/embedcode`,
	}

	// TODO: will be support future below.

	// <script type="text/javascript" src="http://asg.to/js/past_uraui.js" data-title="パ外道親父！"></script><script type="text/javascript" data-title="パち負けた腹いせに息子の嫁を息子の前でる親父！">Purauifla("mcd=X0Bb9lGg4oHcJLjI", 450, 372);</script>

	// <script src="http://static.fc2.com/video/js/outerplayer.min.js" url="http://video.fc2.com/ja/a/content/20150204XER1DTFc/" tk="" tl="【サンプル】 させてみた!!" sj="52000" d="16" w="448" h="380"  suggest="on" charset="UTF-8"></script>

	// <object classid="clsid:d27cdb6e-ae6d-11cf-96b8-444553540000" codebase="http://fpdownload.macromedia.com/pub/shockwave/cabs/flash/swflash.cab#version=8,0,0,0" wmode="transparent" width="448" height="284" id="flv2" align="middle"><param name="allowScriptAccess" value="sameDomain" /><param name="movie" value="http://video.fc2.com/flv2.swf?i=20150206aZzPs9Mg&d=3063&movie_stop=off&no_progressive=1&otag=1&sj=28000&rel=1" /><param name="quality" value="high" /><param name="bgcolor" value="#ffffff" /><param name="allowFullScreen" value="true" /><embed src="http://video.fc2.com/flv2.swf?i=20150206aZzPs9Mg&d=3063&movie_stop=off&no_progressive=1&otag=1&sj=28000&rel=1" quality="high" bgcolor="#ffffff" wmode="transparent" width="448" height="284" name="flv2" align="middle" allowScriptAccess="sameDomain" type="application/x-shockwave-flash" pluginspage="http://www.macromedia.com/go/getflashplayer" allowFullScreen="true" /></object><br /><a href="http://video.fc2.com/content/20150206aZzPs9Mg/" title="対応" rel="nofollow" 【、可愛いが好き</a>

	// <script type="text/javascript" src="http://ero-video.net/js/embed_evplayer.js" data-title="家出少女捕獲』。。。"></script><script type="text/javascript" data-title="出　捕獲』。。。">embedevplayer("mcd=HIzIngBIiNICwg92", 450, 252);</script>

	// <object data="http://embed.redtube.com/player/?id=837335&amp;style=redtube" class="resized" data-title="ギャルのkldalk"><param name="allowfullscreen" value="true"> <param name="AllowScriptAccess" value="always"> <param name="movie" value="http://embed.redtube.com/player/?id=837335&amp;style=redtube"> <param name="FlashVars" value="id=837335&amp;style=redtube&amp;autostart=false"> <embed src="http://embed.redtube.com/player/?id=837335&amp;style=redtube" allowfullscreen="true" allowscriptaccess="always" flashvars="autostart=false" pluginspage="http://www.adobe.com/shockwave/download/download.cgi?P1_Prod_Version=ShockwaveFlash" type="application/x-shockwave-flash" class="resized"></embed></object>

	videoHrefs = []string{
		`www.xvideos.com/video`,
		`jp.xvideos.com/video`,
		`video.fc2.com/content`,
		`video.fc2.com/a/content`,
		`video.fc2.com/ja/a/content`,
		`www.tokyo-tube.com/video`,
		`ero-video.net/movie`,
		`jp.xhamster.com/movies`,
		`www.xhamster.com/movies`,
		`asg.to/contentsPage.html`,
		`www.pornhost.com`,
		`jp.pornhost.com`,
		`www.tube8.com`,
		`www.redtube.com`,
		`www.pornhub.com/view_video.php`,
		`jp.pornhub.com/view_video.php`,
		`japan-whores.com/videos`,
		`www.tokyo-porn-tube.com/video`,
	}
)

type Extractor struct {
	*behavior.UserBehavior

	entry *models.Entry
	doc   *gq.Document
}

func NewExtractor() *Extractor {
	return &Extractor{
		UserBehavior: behavior.NewUserBehavior(),
	}
}

func (e *Extractor) Doc(urlStr string) (*gq.Document, error) {
	resp, err := e.Behave(urlStr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return gq.NewDocumentFromResponse(resp)
}

/*
	Do parsing document
*/
func (e *Extractor) Do(entry *models.Entry) error {
	doc, err := e.Doc(entry.Url)
	if err != nil {
		return err
	}

	e.doc = doc
	e.entry = entry
	return nil
}

/*
	Extracting ExternalURLs
*/
func (e *Extractor) Urls() (urls []string) {
	var re = regexp.MustCompile(strings.Join(videoHrefs, `|`))

	e.doc.Find("a").Each(func(i int, s *gq.Selection) {
		h, ok := s.Attr("href")
		if !ok {
			return
		}
		u, err := urlparse.Parse(h)
		if err != nil || u.Host == "" {
			return
		}

		if re.MatchString(h) {
			urls = append(urls, h)
		}
	})

	return
}

// iframe or script or object
// srcとurlでcontains ableなもの
func (e *Extractor) Codes() (codes []string) {
	var (
		re  = regexp.MustCompile(strings.Join(embedUrls, `|`))
		ok  bool
		val string
	)

	// sel := e.doc.Find("iframe,script,embed,object").FilterFunction(func(i int, s *gq.Selection) bool {
	sel := e.doc.Find("iframe,script").FilterFunction(func(i int, s *gq.Selection) bool {
		val, ok = s.Attr("src")
		if ok && re.MatchString(val) {
			return true
		}
		val, ok = s.Attr("url")
		if ok && re.MatchString(val) {
			return true
		}

		return false
	})

	var buf bytes.Buffer
	sel.Each(func(i int, s *gq.Selection) {
		html.Render(&buf, s.Nodes[0])
		codes = append(codes, str.Clean(buf.String()))
	})

	return
}

/*
	Extracting by string array in Title, Content

	.. note:: 3 chars less word will not extract in function.
*/
func (e *Extractor) ByNames(names []string) (choices []string) {
	for _, name := range names {
		if len([]rune(name)) > 2 && !libm.ReHK3.MatchString(name) {
			if strings.Contains(e.entry.Title, name) {
				choices = append(choices, name)
			} else if strings.Contains(e.entry.Content, name) {
				choices = append(choices, name)
			} else {
				for _, t := range e.entry.Tags {
					if strings.Contains(t.Name, name) {
						choices = append(choices, name)
					}
				}
			}
		}
	}
	return
}

func (e *Extractor) Duration() {}

type Img struct {
	Src string
	Alt string
}

/*
	Extracting img tag
*/
func (e *Extractor) Imgs() (imgs []*Img) {
	u, err := urlparse.Parse(strings.TrimSpace(e.entry.Url))
	if err != nil || u.Host == "" {
		return
	}

	var src string
	domain := strings.Split(u.Host, ":")[0]

	if strings.Contains(domain, "livedoor.jp") {
		src = strings.Split(u.Path, "/")[1]
	} else if strings.Contains(domain, "fc2.com") {
		src = strings.Split(domain, ".")[0]
	} else {
		parts := strings.Split(domain, ".")
		src = parts[len(parts)-2]
	}

	selector := fmt.Sprintf(`img[src*='%s']`, src)
	e.doc.Find(selector).Each(func(i int, s *gq.Selection) {
		src, ok := s.Attr("src")
		if !ok {
			return
		}
		alt, _ := s.Attr("alt")
		if alt == "" {
			alt, _ = s.Attr("title")
		}
		imgs = append(imgs, &Img{Src: src, Alt: alt})
	})

	return
}
