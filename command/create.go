package command

import (
  "github.com/hajarbleh/grafcli/command/create"
  "github.com/urfave/cli"
)

type Create struct {
}

func (c *Create) Commands() []cli.Command {
  dashboard := create.Dashboard{}
  return []cli.Command{
    {
      Name:        "dashboard",
      Usage:       "create dashboard in grafana",
      Action: dashboard.Execute,
    },
  }
}
