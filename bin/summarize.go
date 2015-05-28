package main

import (
	"os"

	_ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	_ "bitbucket.org/ikeikeikeike/antenna/routers"

	"bitbucket.org/ikeikeikeike/antenna/tasks/summarize"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "summarize"
	app.Usage = "Operating the summarize commands"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name: "rssfeed",
			// ShortName: "r",
			Usage: "Summarize Entry model by (RSS)feeds in already registered blogs",
			Action: func(c *cli.Context) {
				err := summarize.RssFeed()
				if err != nil {
					println(err.Error())
				}
			},
		},
		{
			Name: "socialscore",
			// ShortName: "s",
			Usage: "Scoring Entry model by the social share counts",
			Action: func(c *cli.Context) {
				err := summarize.SocialScore()
				if err != nil {
					println(err.Error())
				}
			},
		},
		{
			Name: "inscore",
			// ShortName: "i",
			Usage: "Scoring Entry model by the web server access logs",
			Action: func(c *cli.Context) {
				err := summarize.InScore()
				if err != nil {
					println(err.Error())
				}
			},
		},
		{
			Name: "showcounter",
			// ShortName: "i",
			Usage: "summarize Show page views",
			Action: func(c *cli.Context) {
				err := summarize.Showcounter()
				if err != nil {
					println(err.Error())
				}
			},
		},
		{
			Name: "videoscount",
			// ShortName: "i",
			Usage: "Summarize video starring number for Diva model",
			Action: func(c *cli.Context) {
				err := summarize.DivaVideoscount()
				if err != nil {
					println(err.Error())
				}
			},
		},
		{
			Name: "picturescount",
			// ShortName: "i",
			Usage: "Summarize picture starring number for Character model",
			Action: func(c *cli.Context) {
				err := summarize.CharacterPicturescount()
				if err != nil {
					println(err.Error())
				}
			},
		},
		{
			Name: "animepicturescount",
			// ShortName: "i",
			Usage: "Summarize picture starring number for Anime model",
			Action: func(c *cli.Context) {
				err := summarize.AnimePicturescount()
				if err != nil {
					println(err.Error())
				}
			},
		},
	}

	app.Run(os.Args)
}
