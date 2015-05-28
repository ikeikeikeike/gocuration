package tasks

import (
	"bitbucket.org/ikeikeikeike/antenna/tasks/fillup"
	"bitbucket.org/ikeikeikeike/antenna/tasks/summarize"

	"github.com/astaxie/beego/toolbox"
)

func init() {
	toolbox.AddTask("summarizeRssFeed", summarizeRssFeed())
	toolbox.AddTask("summarizeSocialScore", summarizeSocialScore())
	toolbox.AddTask("summarizeInScore", summarizeInScore())
	toolbox.AddTask("summarizeShowcounter", summarizeShowcounter())
	toolbox.AddTask("summarizeDivaVideosCount", summarizeDivaVideosCount())
	toolbox.AddTask("summarizeCharacterPicturescount", summarizeCharacterPicturescount())
	toolbox.AddTask("summarizeAnimePicturescount", summarizeAnimePicturescount())

	toolbox.AddTask("fillupDivaByApiActresses", fillupDivaByApiActresses())
	toolbox.AddTask("starringDivaImageByGoogleimages", starringDivaImageByGoogleimages())
	toolbox.AddTask("allDivaImageByGoogleimages", allDivaImageByGoogleimages())
	toolbox.AddTask("fillupCharacterByNeoapo", fillupCharacterByNeoapo())
	toolbox.AddTask("fillupCharacterImageByGoogleimages", fillupCharacterImageByGoogleimages())
	toolbox.AddTask("fillupAnimeByNeoapo", fillupAnimeByNeoapo())
	toolbox.AddTask("fillupAnimeImageByGoogleimages", fillupAnimeImageByGoogleimages())

	toolbox.AddTask("starringDivaInfoByWikipedia", starringDivaInfoByWikipedia())
	toolbox.AddTask("allDivaInfoByWikipedia", allDivaInfoByWikipedia())
}

// [seconds] [minutes] [hours] [days] [months] [weeks]

func summarizeRssFeed() *toolbox.Task {
	return toolbox.NewTask("summarizeRssFeed", "0 0 */4 * * *", func() (err error) {
		err = summarize.RssFeed()
		return
	})
}

func summarizeInScore() *toolbox.Task {
	return toolbox.NewTask("summarizeInScore", "0 30 * * * *", func() (err error) {
		err = summarize.InScore()
		return
	})
}

// Execute: xxxx-xx-xx xx:15:00
func summarizeShowcounter() *toolbox.Task {
	return toolbox.NewTask("summarizeShowcounter", "0 15 * * * *", func() (err error) {
		err = summarize.Showcounter()
		return
	})
}

// Execute: xxxx-xx-xx 15:15:00
func summarizeSocialScore() *toolbox.Task {
	return toolbox.NewTask("summarizeSocialScore", "0 15 15 * * *", func() (err error) {
		err = summarize.SocialScore()
		return
	})
}

// Execute: xxxx-xx-xx 02:02:02
func summarizeDivaVideosCount() *toolbox.Task {
	return toolbox.NewTask("summarizeDivaVideosCount", "2 2 2 * * *", func() (err error) {
		err = summarize.DivaVideoscount()
		return
	})
}

// Execute: xxxx-xx-xx 01:01:01
func summarizeCharacterPicturescount() *toolbox.Task {
	return toolbox.NewTask("summarizeCharacterPicturescount", "1 1 1 * * *", func() (err error) {
		err = summarize.CharacterPicturescount()
		return
	})
}

// Execute: xxxx-xx-xx 22:10:10
func summarizeAnimePicturescount() *toolbox.Task {
	return toolbox.NewTask("summarizeAnimePicturescount", "10 10 22 * * *", func() (err error) {
		err = summarize.AnimePicturescount()
		return
	})
}

// Execute: xxxx-xx-03 03:03:03
func fillupCharacterByNeoapo() *toolbox.Task {
	return toolbox.NewTask("fillupCharacterByNeoapo", "3 3 3 3 * *", func() (err error) {
		err = fillup.CharacterByNeoapo()
		return
	})
}

// Execute: xxxx-xx-22 22:22:22
func fillupAnimeByNeoapo() *toolbox.Task {
	return toolbox.NewTask("fillupAnimeByNeoapo", "22 22 22 22 * *", func() (err error) {
		err = fillup.AnimeByNeoapo()
		return
	})
}

// Execute: xxxx-xx-04 04:04:04
func fillupDivaByApiActresses() *toolbox.Task {
	return toolbox.NewTask("fillupDivaByApiActresses", "4 4 4 4 * *", func() (err error) {
		err = fillup.DivaByApiActresses()
		return
	})
}

// Execute: xxxx-xx-xx 06:06:06
func starringDivaImageByGoogleimages() *toolbox.Task {
	return toolbox.NewTask("starringDivaImageByGoogleimages", "6 6 6 * * *", func() (err error) {
		err = fillup.StarringDivaImageByGoogleimages()
		return
	})
}

// Execute: xxxx-1,3,5,7,,,-15 00:00:00
func allDivaImageByGoogleimages() *toolbox.Task {
	return toolbox.NewTask("allDivaImageByGoogleimages", "0 0 0 15 */2 *", func() (err error) {
		err = fillup.AllDivaImageByGoogleimages()
		return
	})
}

// Execute: xxxx-xx-xx 05:05:05
func fillupCharacterImageByGoogleimages() *toolbox.Task {
	return toolbox.NewTask("fillupCharacterImageByGoogleimages", "5 5 5 * * *", func() (err error) {
		err = fillup.CharacterImageByGoogleimages()
		return
	})
}

// Execute: xxxx-xx-xx 23:10:10
func fillupAnimeImageByGoogleimages() *toolbox.Task {
	return toolbox.NewTask("fillupAnimeImageByGoogleimages", "10 10 23 * * *", func() (err error) {
		err = fillup.AnimeImageByGoogleimages()
		return
	})
}

// Execute: xxxx-xx-xx 07:07:07
func starringDivaInfoByWikipedia() *toolbox.Task {
	return toolbox.NewTask("starringDivaInfoByWikipedia", "7 7 7 * * *", func() (err error) {
		err = fillup.StarringDivaInfoByWikipedia()
		return
	})
}

// Execute: xxxx-1,4,7,,,-07 07:07:07
func allDivaInfoByWikipedia() *toolbox.Task {
	return toolbox.NewTask("allDivaInfoByWikipedia", "7 7 7 7 */3 *", func() (err error) {
		err = fillup.AllDivaInfoByWikipedia()
		return
	})
}
