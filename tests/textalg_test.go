package main

import (
	"fmt"
	"strings"
	"testing"

	t2 "github.com/arbovm/levenshtein"
	"github.com/k0kubun/pp"
	t1 "github.com/masatana/go-textdistance"
	t3 "github.com/texttheater/golang-levenshtein/levenshtein"
)

type TestDistanceSrcs struct {
	s1 string
	s2 string
}

func GetTestDistanceSrcs() []TestDistanceSrcs {
	return []TestDistanceSrcs{
		{s1: "doujinsieromangasouko", s2: "http://static.fc2.com/image/headbar/sh_fc2blogheadbar_logo.png"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/050505963.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://pamera-buy.com/afl/data.php?i=528850e5c3afc&m=52d641f0bad7d"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-64.fc2.com/d/o/u/doujinsieromanga18/hi01_R.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-64.fc2.com/d/o/u/doujinsieromanga18/001_R_201408251600385bb.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-64.fc2.com/d/o/u/doujinsieromanga18/001_R_20140825155412498.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-64.fc2.com/d/o/u/doujinsieromanga18/Rebecca1_R.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-64.fc2.com/d/o/u/doujinsieromanga18/001_R_20140816144245043.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-64.fc2.com/d/o/u/doujinsieromanga18/001_R_20140711152436e7a.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-64.fc2.com/d/o/u/doujinsieromanga18/001_R_20140812144935d28.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-64.fc2.com/d/o/u/doujinsieromanga18/001_R_20140819160005d95.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-64.fc2.com/d/o/u/doujinsieromanga18/045_R_20140826144050dcb.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-64.fc2.com/d/o/u/doujinsieromanga18/001_R_201408191527411d0.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-64.fc2.com/d/o/u/doujinsieromanga18/index_01_1_R.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-64.fc2.com/d/o/u/doujinsieromanga18/2014082517104127b.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-64.fc2.com/d/o/u/doujinsieromanga18/01_R_20140819145119574.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-64.fc2.com/d/o/u/doujinsieromanga18/01_R_2014082516284078a.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-64.fc2.com/d/o/u/doujinsieromanga18/001_R_20140819142735478.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/050505963.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-64.fc2.com/d/o/u/doujinsieromanga18/85.gif"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/00001_R_20150520142337264.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/050505963.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/00002_R_20150520142339c7d.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/00003_R_201505201423394c0.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/00004_R_20150520142342cdd.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/00005_R_201505201423424f3.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/00006_R_20150520142400a1b.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/00007_R_201505201424016a9.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/00008_R_201505201424020f0.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/00009_R_201505201424046c4.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/00010_R_201505201424053de.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/00011_R_20150520142422c5e.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/00012_R_201505201424245b3.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/00013_R_201505201424255fb.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/00013_R_201505201424255fb.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/00014_R_201505201424275ea.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/00015_R_20150520142428176.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/00016_R_20150520142447421.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/00017_R_20150520142449579.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://static.fc2.com/image/clap/number/orange/1.gif"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/050505963.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/c/h/i/chijodougasokuhou/blog-entry-2346.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://livedoor.blogimg.jp/girl_4s/imgs/b/d/bd5c13bf.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-72.fc2.com/c/h/i/chijodougasokuhou/blog-entry-2328.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-74.fc2.com/e/r/o/eromaniaxmovie/_thumb_20150530161105.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-75.fc2.com/m/o/g/mogiero/52320041506201539.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/c/h/i/chijodougasokuhou/blog-entry-2346.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://livedoor.blogimg.jp/girl_4s/imgs/b/d/bd5c13bf.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-72.fc2.com/c/h/i/chijodougasokuhou/blog-entry-2328.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-74.fc2.com/e/r/o/eromaniaxmovie/_thumb_20150530161105.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-75.fc2.com/m/o/g/mogiero/52320041506201539.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-64.fc2.com/d/o/u/doujinsieromanga18/20140822110106964.gif"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-64.fc2.com/d/o/u/doujinsieromanga18/82.png"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-64.fc2.com/d/o/u/doujinsieromanga18/34301166.png"},
		{s1: "doujinsieromangasouko", s2: "http://blogranking.fc2.com/ranking_banner/a_04.gif"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-50.fc2.com/d/o/u/doujinsieromanga18/path4146.png"},
		{s1: "doujinsieromangasouko", s2: "http://www.trackword.biz/img/minilogov.gif"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/001_R_2015020516224920f.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-64.fc2.com/d/o/u/doujinsieromanga18/001_R_201407161621252d6.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/001_R_20150121150331703.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-64.fc2.com/d/o/u/doujinsieromanga18/001_R_20140819160005d95.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/00001_R_20150224143236fa0.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/001_R_20150113165937543.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-64.fc2.com/d/o/u/doujinsieromanga18/201409241455280c1.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-64.fc2.com/d/o/u/doujinsieromanga18/001_R_2014091614453105b.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/001_R_20150317153543b7b.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-64.fc2.com/d/o/u/doujinsieromanga18/001_R_20140717155041cdd.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/00001_R_20150210152415dad.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-64.fc2.com/d/o/u/doujinsieromanga18/001_R_201407241537353e1.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-64.fc2.com/d/o/u/doujinsieromanga18/001_R_20140819143230686.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://img.i2i.jp/sr/ad/amz2.gif"},
		{s1: "doujinsieromangasouko", s2: "http://twitbtn.com/images/buttons/button_new24a.png"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/g1.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/20150204164948ce8.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/20150204164951673.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/20150204164953d00.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/20150204165005603.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/20150204165008582.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-73.fc2.com/d/o/u/doujinsieromanga18/20150204165009344.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-64.fc2.com/d/o/u/doujinsieromanga18/729d7e422.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://rranking8.ziyu.net/rranking.gif"},
		{s1: "doujinsieromangasouko", s2: "http://analyzer5.fc2cn.com/ana/icon0.gif"},
		{s1: "doujinsieromangasouko", s2: "doujinsieromangasouko"},
		{s1: "doujinsieromangasouko", s2: "doujinsieromanga18"},
		{s1: "doujinsieromangasouko", s2: "/d/o/u/doujinsieromanga18/729d7e422.jpg"},
		{s1: "doujinsieromangasouko", s2: "http://blog-imgs-64.fc2.com/d/o/u/doujinsieromanga18/729d7e422.jpg"},
	}
}

func TestLevenshtein1(t *testing.T) {

	for _, src := range GetTestDistanceSrcs() {
		s1, s2 := src.s1, src.s2

		fmt.Printf("%v -> %v\n", s2, t1.LevenshteinDistance(s1, s2))
		fmt.Printf("%v -> %v\n", s2, t1.DamerauLevenshteinDistance(s1, s2))
		a, b := t1.JaroDistance(s1, s2)
		fmt.Printf("%v -> %v:%v\n", s2, a, b)
		fmt.Printf("%v -> %v\n", s2, t1.JaroWinklerDistance(s1, s2))
	}
}

func TestLevenshtein2(t *testing.T) {
	for _, src := range GetTestDistanceSrcs() {
		s1, s2 := src.s1, src.s2

		fmt.Printf("The distance between '%v' and '%v' is %v\n", s1, s2, t2.Distance(s1, s2))
	}
}

func TestLevenshtein3(t *testing.T) {
	for _, src := range GetTestDistanceSrcs() {
		s1, s2 := src.s1, src.s2

		fmt.Println(t3.DistanceForStrings([]rune(s1), []rune(s2), t3.DefaultOptions))
	}

}

func TestLevenshtein4(t *testing.T) {
	for _, src := range GetTestDistanceSrcs() {
		same := false
		s1, s2 := src.s1, src.s2

		for _, term := range strings.Split(s2, "/") {
			if t1.LevenshteinDistance(s1, term) < 10 {
				same = true
			}
		}

		pp.Println(s1, s2, same)
		println("")
	}
}
