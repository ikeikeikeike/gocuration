package main

import (
	"testing"

	"github.com/k0kubun/pp"

	_ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	"bitbucket.org/ikeikeikeike/antenna/models/site"
	_ "bitbucket.org/ikeikeikeike/antenna/routers"
)

func TestRegex(t *testing.T) {

	pp.Println("com", site.SelectDomains("com"))
	pp.Println("xvideos", site.SelectDomains("xvideos"))
	pp.Println("xvideos.com", site.SelectDomains("xvideos.com"))
	pp.Println("www.xvideos.com", site.SelectDomains("www.xvideos.com"))
	pp.Println("http://www.xvideos.com", site.SelectDomains("http://www.xvideos.com"))
	pp.Println("http://flashservice.xvideos.com/embedframe?any=12", site.SelectDomains("http://flashservice.xvideos.com/embedframe?any=12"))

	pp.Println("xvideos", site.SelectDomains(`<iframe src="http://flashservice.xvideos.com/embedframe/3224660" frameborder="0" width="510" height="400" scrolling="no"></iframe>`))
	pp.Println("fc2", site.SelectDomains(`<script src="http://static.fc2.com/video/js/outerplayer.min.js" url="http://video.fc2.com/ja/a/content/20150207ZrkqfcMR/" tk="" tl="" sj="69000" d="12" w="448" h="288"  suggest="on" charset="UTF-8"></script`))
	pp.Println("tokyo-tube", site.SelectDomains(`<script type="text/javascript" src="http://www.tokyo-tube.com/embedcode/v149399/u/player/w452/h361" data-title=""></script>`))
	pp.Println("ero-video", site.SelectDomains(`<script type="text/javascript" src="http://ero-video.net/js/embed_evplayer.js" data-title=""></script><script type="text/javascript" data-title="">embedevplayer("mcd=KSumMQzBi84HCsof", 450, 341);</script>`))
	pp.Println("xhamster", site.SelectDomains(`<iframe width="510" height="400" src="http://xhamster.com/xembed.php?video=2594535" frameborder="0" scrolling="no" data-title=""></iframe>`))
	pp.Println("japan-whores", site.SelectDomains(`<iframe src="http://japan-whores.com/embed/161739/" width="640" height="480" frameborder="0" scrolling="no" allowfullscreen="" webkitallowfullscreen="" mozallowfullscreen="" oallowfullscreen="" msallowfullscreen="">`))

	pp.Println("blank", site.SelectDomains(""))

	// if len(domains) <= 0 { t.Error(err) }
}
