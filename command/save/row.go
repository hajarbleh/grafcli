package save

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/hajarbleh/grafcli/client"
	"github.com/hajarbleh/grafcli/config"
	"github.com/hajarbleh/grafcli/template"
	"github.com/urfave/cli"
)

type Row struct {
	Filename      string
	DashboardName string
}

func (r *Row) Execute(ctx *cli.Context) error {
	if r.Filename == "" {
		fmt.Println(errors.New("Error: Required flag \"filename\"(-f) are not set!"))
		return errors.New("Error: Required flag \"filename\"(-f) are not set!")
	}

	if r.DashboardName == "" {
		fmt.Println(errors.New("Error: Required flag \"dashboard name\"(-d) are not set!"))
		return errors.New("Error: Required flag \"dashboard name\"(-d) are not set!")
	}

	data, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		fmt.Sprintf("Fatal error loading file: %s \n", err)
		return errors.New(fmt.Sprintf("Fatal error loading file: %s \n", err))
	}

	var newRow template.DashboardRow
	_ = yaml.Unmarshal([]byte(data), &newRow)

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

	for idx, d := range dashboardExtended.Dashboard.Rows {
		if strings.ToLower(d.Title) == strings.ToLower(newRow.Title) {
			dashboardExtended.Dashboard.Rows[idx] = newRow
			break
		}
	}

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
			Name:        "f",
			Usage:       "filename to save the resource",
			Value:       "",
			Destination: &r.Filename,
		},
		cli.StringFlag{
			Name:        "d",
			Usage:       "dashboard URI to save the row",
			Value:       "",
			Destination: &r.DashboardName,
		},
	}
}
