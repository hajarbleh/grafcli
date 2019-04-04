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

type Row struct {
	DashboardName string
}

func (r *Row) Execute(ctx *cli.Context) error {
	rName := ctx.Args().Get(0)
	if rName == "" {
		fmt.Println("must specify row name")
		return errors.New("must specify row name")
	}

	c, err := config.Read()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	grafana := client.NewGrafana(c.Url, c.ApiKey)
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

	newRow := template.NewDashboardRow(rName)
	dashboardExtended.Dashboard.Rows = append(dashboardExtended.Dashboard.Rows, newRow)

	jsonDashboard, _ := json.Marshal(dashboardExtended.Dashboard)
	if _, err := grafana.CreateDashboard(string(jsonDashboard), false, "Updated by grafcli"); err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Dashboard successfully saved!")
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
