package command

import (
  "github.com/hajarbleh/grafcli/command/get"
  "github.com/urfave/cli"
)

type Get struct {
}

func (g *Get) Commands() []cli.Command {
  dashboard := get.Dashboard{}
  return []cli.Command{
    {
      Name:   "dashboard",
      Usage:  "get dashboard json",
      Subcommands: dashboard.Commands(),
    },
  }
}
