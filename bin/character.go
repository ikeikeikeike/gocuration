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

	app.Name = "character"
	app.Usage = "Opps to the Character model"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:  "push",
			Usage: "Add a got prefix(page_number) records from Neoapo into the Character model",
			Action: func(c *cli.Context) {
				gyou := c.Args().First()
				if gyou == "" {
					println("Error: Does not multiple args. Set more than one args[prefix(page_number)]")
					return
				}

				err := fillup.CharacterByNeoapoByPrefix(c.Args().First())
				if err != nil {
					println(err.Error())
				}
			},
		},
		{
			Name:  "fillup",
			Usage: "Fill it up to the Character model",
			Action: func(c *cli.Context) {
				err := fillup.CharacterByNeoapo()
				if err != nil {
					println(err.Error())
				}
			},
		},
		{
			Name: "googleimage",
			// ShortName: "g",
			Usage: "Fill image up to the Character model by google image api",
			Action: func(c *cli.Context) {
				err := fillup.CharacterImageByGoogleimages()
				if err != nil {
					println(err.Error())
				}
			},
		},
		// {
		// Name: "wikipedia",
		// // ShortName: "g",
		// Usage: "Fill infomation up to the diva model by wikimedia api",
		// Action: func(c *cli.Context) {
		// err := fillup.DivaInfoByWikipedia()
		// if err != nil {
		// println(err.Error())
		// }
		// },
		// },
	}

	app.Run(os.Args)
}
