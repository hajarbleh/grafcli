package command

import (
	"github.com/hajarbleh/grafcli/command/describe"
	"github.com/urfave/cli"
)

type Describe struct {
}

func (d *Describe) Commands() []cli.Command {
	dashboard := describe.Dashboard{}
	row := describe.Row{}

	return []cli.Command{
		{
			Name:   "dashboard",
			Usage:  "get dashboard in yaml format",
			Action: dashboard.Execute,
		},
		{
			Name:   "row",
			Usage:  "get row in yaml format",
			Flags:  row.Flags(),
			Action: row.Execute,
		},
	}
}
