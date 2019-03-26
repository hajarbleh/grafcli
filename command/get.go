package command

import (
	"github.com/hajarbleh/grafcli/command/get"
	"github.com/urfave/cli"
)

type Get struct {
}

func (g *Get) Commands() []cli.Command {
	dashboard := get.Dashboard{}
	rows := get.Rows{}
	return []cli.Command{
		{
			// refactor this command in next major
			Name:        "dashboard",
			Usage:       "get dashboard",
			Subcommands: dashboard.Commands(),
		},
		{
			Name:        "rows",
			Usage:       "get row list",
			Flags: rows.Flags(),
			Action: rows.Execute,
		},
	}
}
