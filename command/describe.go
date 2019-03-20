package command

import (
  "github.com/hajarbleh/grafcli/command/describe"
  "github.com/urfave/cli"
)

type Describe struct {
}

func (d *Describe) Commands() []cli.Command {
  dashboard := describe.Dashboard{}
  return []cli.Command{
    {
      Name:   "dashboard",
      Usage:  "get dashboard json",
      Action: dashboard.Execute,
    },
  }
}
