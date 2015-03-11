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

	app.Name = "diva"
	app.Usage = "Opps to the Diva model"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name: "push",
			// ShortName: "p",
			Usage: "Add a got prefix(gyou) records from apisource into the diva model",
			Action: func(c *cli.Context) {
				gyou := c.Args().First()
				if gyou == "" {
					println("Error: Does not multiple args. Set more than one args[prefix(gyou)]")
					return
				}

				err := fillup.DivaByApiActress(c.Args().First())
				if err != nil {
					println(err.Error())
				}
			},
		},
		{
			Name: "fillup",
			// ShortName: "u",
			Usage: "Fill it up to the diva model",
			Action: func(c *cli.Context) {
				err := fillup.DivaByApiActresses()
				if err != nil {
					println(err.Error())
				}
			},
		},
		{
			Name: "googleimage",
			// ShortName: "g",
			Usage: "Fill image up to the diva model by google image api",
			Action: func(c *cli.Context) {
				err := fillup.DivaImageByGoogleimages()
				if err != nil {
					println(err.Error())
				}
			},
		},
		{
			Name: "wikipedia",
			// ShortName: "g",
			Usage: "Fill infomation up to the diva model by wikimedia api",
			Action: func(c *cli.Context) {
				err := fillup.DivaInfoByWikipedia()
				if err != nil {
					println(err.Error())
				}
			},
		},
	}

	app.Run(os.Args)
}
