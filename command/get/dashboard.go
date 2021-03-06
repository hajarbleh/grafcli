package get

import (
	dashboard "github.com/hajarbleh/grafcli/command/get/dashboard"
	"github.com/urfave/cli"
)

type Dashboard struct {
}

func (d *Dashboard) Commands() []cli.Command {
	list := dashboard.List{}
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "get list of dashboard in your app",
			Action: list.Execute,
		},
	}
}
