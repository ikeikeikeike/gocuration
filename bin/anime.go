package main

import (
	"os"

	_ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	_ "bitbucket.org/ikeikeikeike/antenna/routers"

	"bitbucket.org/ikeikeikeike/antenna/tasks/fillup"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "anime"
	app.Usage = "Opps to the Anime model"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:  "fromentries",
			Usage: "Fill up(Summarize) Picture.Anime model from Entry model",
			Action: func(c *cli.Context) {
				err := fillup.AnimeFromEntries()
				if err != nil {
					println(err.Error())
				}
			},
		},
		{
			Name:  "fromcharacter",
			Usage: "Fill up Anime model from Character model",
			Action: func(c *cli.Context) {
				err := fillup.AnimeByCharacterModel()
				if err != nil {
					println(err.Error())
				}
			},
		},
		{
			Name:  "push",
			Usage: "Add a got prefix(page_number) records from Neoapo into the Anime model",
			Action: func(c *cli.Context) {
				gyou := c.Args().First()
				if gyou == "" {
					println("Error: Does not multiple args. Set more than one args[prefix(page_number)]")
					return
				}

				err := fillup.AnimeByNeoapoByPrefix(c.Args().First())
				if err != nil {
					println(err.Error())
				}
			},
		},
		{
			Name:  "fillup",
			Usage: "Fill it up to the Anime model",
			Action: func(c *cli.Context) {
				err := fillup.AnimeByNeoapo()
				if err != nil {
					println(err.Error())
				}
			},
		},
		// {
		// Name:  "fillupkana",
		// Usage: "fillup to katakana in model",
		// Action: func(c *cli.Context) {
		// option := c.Args().First()
		// if option == "" {
		// println("Error: Does not multiple args. Set more than one args[mecab option]. e.g. /dic/mecab-ipadic-neologd")
		// return
		// }

		// err := fillup.AnimeKanaByMecab(option)
		// if err != nil {
		// println(err.Error())
		// }
		// },
		// },
		{
			Name: "googleimage",
			// ShortName: "g",
			Usage: "Fill image up to the Anime model by google image api",
			Action: func(c *cli.Context) {
				err := fillup.AnimeImageByGoogleimages()
				if err != nil {
					println(err.Error())
				}
			},
		},
		// {
		// Name: "wikipedia",
		// // ShortName: "g",
		// Usage: "Fill infomation up to the Anime model by wikimedia api",
		// Action: func(c *cli.Context) {
		// err := fillup.AnimeInfoByWikipedia()
		// if err != nil {
		// println(err.Error())
		// }
		// },
		// },
	}

	app.Run(os.Args)
}
