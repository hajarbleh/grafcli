package describe

import (
	"errors"
	"fmt"

	"github.com/ghodss/yaml"
	"github.com/hajarbleh/grafcli/client"
	"github.com/hajarbleh/grafcli/config"
	"github.com/urfave/cli"
)

type Dashboard struct {
}

func (d *Dashboard) Execute(ctx *cli.Context) error {
	dName := ctx.Args().Get(0)
	if dName == "" {
		fmt.Println("must specify dashboard name")
		return errors.New("must specify dashboard name")
	}

	c, err := config.Read()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	grafana := client.NewGrafana(c.Url, c.ApiKey)
	body, err := grafana.GetDashboard(dName)
	if err != nil {
		fmt.Println(err)
		return err
	}

	y, err := yaml.JSONToYAML([]byte(body))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string([]byte(y)))

	return nil
}

func (d *Dashboard) Flags() []cli.Flag {
	return []cli.Flag{}
}
