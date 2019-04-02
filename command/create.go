package command

import (
	"github.com/hajarbleh/grafcli/command/create"
	"github.com/urfave/cli"
)

type Create struct {
}

func (c *Create) Commands() []cli.Command {
	dashboard := create.Dashboard{}
	row := create.Row{}
	panel := create.Panel{}
	return []cli.Command{
		{
			Name:   "dashboard",
			Usage:  "create dashboard in grafana",
			Action: dashboard.Execute,
		},
		{
			Name:   "row",
			Usage:  "create row in grafana dashboard (will be inserted in bottom row)",
			Action: row.Execute,
			Flags:  row.Flags(),
		},
		{
			Name:   "panel",
			Usage:  "create panel in grafana dashboard",
			Action: panel.Execute,
			Flags:  panel.Flags(),
		},
	}
}
