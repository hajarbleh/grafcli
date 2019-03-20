package main

import (
	"fmt"
	"os"

	"github.com/hajarbleh/grafcli/command"
	"github.com/urfave/cli"
)

func main() {
	set := &command.Set{}
  get := &command.Get{}
  describe := &command.Describe{}
  save := &command.Save{}

	app := cli.NewApp()
	app.Name = "grafcli"
	app.Version = "0.0.0"
	app.Usage = "Maintain your grafana dashboards"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "John Stephanus Peter",
			Email: "johnstephanus@ymail.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "set",
			Usage:  "set context entry for specified context name",
			Action: set.Execute,
			Flags:  set.Flags(),
		},
		{
		  Name: "get",
		  Usage: "fetch resources from grafana",
		  Subcommands: get.Commands(),
		},
    {
      Name: "describe",
      Usage: "describe resources from grafana",
      Subcommands: describe.Commands(),
    },
    {
      Name: "save",
      Usage: "save resource to grafana",
      Subcommands: save.Commands(),
    },
	}

	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Println("unknown command %s", command)
	}

	app.Run(os.Args)
}
