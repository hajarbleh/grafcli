package command

import (
  "github.com/hajarbleh/grafcli/command/describe"
  "github.com/urfave/cli"
)

type Describe struct {
}

func (d *Describe) Commands() []cli.Command {
  dashboard := describe.Dashboard{}
  panel := describe.Panel{}

  return []cli.Command{
    {
      Name:   "dashboard",
      Usage:  "get dashboard json",
      Action: dashboard.Execute,
    },
    {
      Name: "panel",
      Usage: "get panel json of specified dashboard",
      Action: panel.Execute,
      Flags: panel.Flags(),
    },
  }
}
