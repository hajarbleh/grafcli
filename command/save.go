package command

import (
	"github.com/hajarbleh/grafcli/command/save"
	"github.com/urfave/cli"
)

type Save struct {
}

func (s *Save) Commands() []cli.Command {
	dashboard := save.Dashboard{}
	row := save.Row{}
	return []cli.Command{
		{
			Name:   "dashboard",
			Usage:  "save dashboard from file",
			Action: dashboard.Execute,
			Flags:  dashboard.Flags(),
		},
		{
			Name: "row",
			Usage: "save row from file",
			Action: row.Execute,
			Flags: row.Flags(),
		},
	}
}
