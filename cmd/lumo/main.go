package main

import (
	"fmt"
	"os"

	"github.com/rg-devel/lumo"
	"github.com/urfave/cli"
	log "github.com/Sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000",
		FullTimestamp: true,
	})
}

func main() {

	log.SetOutput(os.Stdout)

	log.WithFields(log.Fields{
		"string field": "string value",
		"number": 1,
	}).Info("starting lumo")
	log.Warn("warning message without fields")

	app := cli.NewApp()

	app.Name = "lumo"
	app.Version = "0.0.1"
	app.Usage = "Log parser in Go"
	app.Commands = []cli.Command{
		{
			Name:        "list",
			Aliases:     []string{"l"},
			Usage:       "list jobs",
			Description: "List jobs filtered by the arguments",
			Subcommands: []cli.Command{
				{
					Name:        "all",
					Aliases:     []string{"a"},
					Usage:       "list all jobs",
					Description: "List all jobs in the log(s)",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name",
							Value: "Bob",
							Usage: "Name of the person to greet",
						},
					},
					Action: func(c *cli.Context) error {
						fmt.Println("Hello,", c.String("name"))
						return nil
					},
				},
			},
		},{
			Name:        "show",
			Usage:       "show details",
			Description: "Show details of the arguments",
			Subcommands: []cli.Command{
				{
					Name:        "job",
					Aliases:     []string{"a"},
					Usage:       "show all entries of a given job",
					Description: "List all entries related to a given job",
					Flags: []cli.Flag{
						cli.IntFlag{
							Name:  "id",
							Value: 0,
							Usage: "Job ID",
						},
					},
					Action: func(c *cli.Context) error {
						fmt.Printf("This is where all the entries for job ID %d is displayed", c.Int("id"))
						return nil
					},
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Printf("the next example")
				return nil
			},
		},
	}
	app.Action = func(c *cli.Context) error {
		fmt.Printf("Hello %q\n", c.Args().Get(0))
		err := lumo.Foo()
		if err != nil {
			return err
		}
		fmt.Println("done")
		return nil
	}
	_ = app.Run(os.Args)
}
