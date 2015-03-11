package main

import (
	"fmt"
	"testing"

	"github.com/astaxie/beego"
	"github.com/ikeikeikeike/go-googleimages"
	"github.com/k0kubun/pp"

	_ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	_ "bitbucket.org/ikeikeikeike/antenna/routers"
)

func TestImgExtract(t *testing.T) {
	c := googleimages.NewClient()
	c.Header("User-Agent", beego.AppConfig.String("UserAgent"))
	result, _ := c.Fetch(fmt.Sprintf("%s 女優", "みどり"))
	pp.Println(result)
}

// func TestImageOk(t *testing.T) {
// var img string

// img = "http://dafadsf.dcom/dafas.jpg"
// info, err := image.NewFileInfo(img)
// pp.Println(info, err)

// img = "http://image.eiga.k-img.com/images/special/1742/635.jpg%3F1334735359"
// info, err = image.NewFileInfo(img)
// pp.Println(info, err)

// img = "http://image.eiga.k-img.com/images/special/1742/635.jpg"
// info, err = image.NewFileInfo(img)
// pp.Println(info, err)

// img = "http://cdn.mkimg.carview.co.jp/minkara/photo/000/003/604/156/3604156/p3.jpg%3Fct%3Dd4ea00ea0cec"
// info, err = image.NewFileInfo(img)
// pp.Println(info, err)
// }

// func TestDivaExtract(t *testing.T) {

// e := &models.Entry{Id: 28}
// e.Read()
// e.LoadRelated()

// ext := video.NewExtractor()
// ext.Header("User-Agent", beego.AppConfig.String("UserAgent"))

// if err := ext.Do(e); err != nil {
// t.Fatalf("Error Entry{id=%s}: %s", e.Id, err)
// }

// // Save Diva
// var divas []*models.Diva
// for _, name := range ext.ByNames(diva.DivasName()) {
// d := &models.Diva{Name: name}
// d.Read("Name")
// divas = append(divas, d)
// }

// if false {
// t.Fatalf("Rss fetch error: %s", "")
// }

// }
