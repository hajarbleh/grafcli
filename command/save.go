package command

import (
	"github.com/hajarbleh/grafcli/command/save"
	"github.com/urfave/cli"
)

type Save struct {
}

func (s *Save) Commands() []cli.Command {
	dashboard := save.Dashboard{}
	return []cli.Command{
		{
			Name:   "dashboard",
			Usage:  "save dashboard from file",
			Action: dashboard.Execute,
			Flags:  dashboard.Flags(),
		},
	}
}
