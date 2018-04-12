package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	var recusive bool
	var language string
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "recusive,r",
			Usage:       "recusive or not",
			Destination: &recusive,
		},
		cli.StringFlag{
			Name:        "language,l",
			Value:       "Chinese",
			Usage:       "choose the language",
			Destination: &language,
		},
	}
	app.Action = func(c *cli.Context) {
		var cmd string
		if c.NArg() > 0 {
			cmd = c.Args()[0]
			fmt.Printf("cmd filePaht:%s\t", cmd)
		}
		fmt.Printf("language:%s\t", language)
		fmt.Printf("recusive:%v\t", recusive)

	}
	/*app.Action = func(c *cli.Context) error {
		var cmd string
		if c.NArg() > 0 {
			cmd = c.Args()[0]
			fmt.Println("cmd is ", cmd)
		}
		fmt.Println("recusive is ", recusive)
		fmt.Println("language is ", language)
		return nil
	}*/
	app.Run(os.Args)
}
