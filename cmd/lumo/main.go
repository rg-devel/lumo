package main

import (
	"fmt"
	"os"

	"github.com/rg-devel/lumo"
	"github.com/urfave/cli"
)

var version = "master"

func main() {
	app := cli.NewApp()
	app.Name = "lumo"
	app.Version = version
	app.Author = "Carlos Alexandro Becker (root@carlosbecker.com)"
	app.Usage = "This is an lumo app written in Go"
	app.Action = func(c *cli.Context) error {
		err := lumo.Foo()
		if err != nil {
			return err
		}
		fmt.Println("Done!")
		return nil
	}
	_ = app.Run(os.Args)
}
