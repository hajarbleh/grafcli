package describe

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/hajarbleh/grafcli/client"
	"github.com/hajarbleh/grafcli/config"
	"github.com/hajarbleh/grafcli/template"
	"github.com/urfave/cli"
	"github.com/ghodss/yaml"
)

type Row struct {
	DashboardName string
}

func (r *Row) Execute(ctx *cli.Context) error {
	rName := ctx.Args().Get(0)
	if rName == "" {
		fmt.Println("must specify row name")
		return errors.New("must specify row name")
	}

	grafana := client.NewGrafana(config.URL, config.APIKey)
	body, err := grafana.GetDashboard(r.DashboardName)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var dashboardExtended template.DashboardExtended
	err = json.Unmarshal([]byte(body), &dashboardExtended)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	dRow := dashboardExtended.Dashboard.Rows

	for _, row := range dRow {
		if strings.ToLower(row.Title) == strings.ToLower(rName) {
			out, _ := yaml.Marshal(&row)
			fmt.Println(string([]byte(out)))
			return nil
		}
	}

	fmt.Println("Dashboard row not found")

	return nil
}

func (r *Row) Flags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "d",
			Usage:       "specify dashboard name",
			Value:       "",
			Destination: &r.DashboardName,
		},
	}
}
