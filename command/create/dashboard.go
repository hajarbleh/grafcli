package create

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hajarbleh/grafcli/client"
	"github.com/hajarbleh/grafcli/config"
	"github.com/hajarbleh/grafcli/template"
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

	jsonDashboard, _ := json.Marshal(template.NewDashboard(dName))
	grafana := client.NewGrafana(config.URL, config.APIKey)
	body, err := grafana.CreateDashboard(string(jsonDashboard), false, "Updated by grafcli")
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(string([]byte(body)))
	return nil
}

func (d *Dashboard) Flags() []cli.Flag {
	return []cli.Flag{}
}
