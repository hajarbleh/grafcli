package get

import (
	"fmt"
	dashboard "github.com/hajarbleh/grafcli/command/get/dashboard"
	"github.com/urfave/cli"
)

type Dashboard struct {
}

func (d *Dashboard) Execute(ctx *cli.Context) error {
	fmt.Println("Example executionsss")
	return nil
}

func (d *Dashboard) Flags() []cli.Flag {
	return []cli.Flag{}
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
