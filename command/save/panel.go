package save

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/hajarbleh/grafcli/client"
	"github.com/hajarbleh/grafcli/config"
	"github.com/hajarbleh/grafcli/template"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

type Panel struct {
	Filename      string
	DashboardName string
}

func (p *Panel) Execute(ctx *cli.Context) error {
	if p.Filename == "" {
		fmt.Println(errors.New("Error: Required flag \"filename\"(-f) are not set!"))
		return errors.New("Error: Required flag \"filename\"(-f) are not set!")
	}

	if p.DashboardName == "" {
		fmt.Println(errors.New("Error: Required flag \"dashboard name\"(-d) are not set!"))
		return errors.New("Error: Required flag \"dashboard name\"(-d) are not set!")
	}

	data, err := ioutil.ReadFile(p.Filename)
	if err != nil {
		fmt.Sprintf("Fatal error loading file: %s \n", err)
		return errors.New(fmt.Sprintf("Fatal error loading file: %s \n", err))
	}

	var newPanel map[string]interface{}
	_ = yaml.Unmarshal([]byte(data), &newPanel)

	grafana := client.NewGrafana(config.URL, config.APIKey)
	body, err := grafana.GetDashboard(p.DashboardName)
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

	for ridx, r := range dashboardExtended.Dashboard.Rows {
		for pidx, p := range r.Panels {
			if p["id"] == newPanel["id"] {
				dashboardExtended.Dashboard.Rows[ridx].Panels[pidx] = newPanel
			}
		}
	}

	jsonDashboard, _ := json.Marshal(dashboardExtended.Dashboard)
	if _, err := grafana.CreateDashboard(string(jsonDashboard), false, "Updated by grafcli"); err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Panel successfully saved!")
	return nil
}

func (p *Panel) Flags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "f",
			Usage:       "filename to save the resource",
			Value:       "",
			Destination: &p.Filename,
		},
		cli.StringFlag{
			Name:        "d",
			Usage:       "dashboard URI to save the panel",
			Value:       "",
			Destination: &p.DashboardName,
		},
	}
}
