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

	app.Name = "oneshot"
	app.Usage = "oneshot commands"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name: "fillup_imagecount_on_picture",
			// ShortName: "r",
			Usage: "fill up image count picture",
			Action: func(c *cli.Context) {
				err := fillup.PictureImageCount()
				if err != nil {
					println(err.Error())
				}
			},
		},
	}

	app.Run(os.Args)
}
