package describe

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/hajarbleh/grafcli/config"
	"github.com/hajarbleh/grafcli/template"
	"github.com/urfave/cli"
)

type Panel struct {
	DashboardName string
	RowName       string
	PanelName     string
}

func (p *Panel) Execute(ctx *cli.Context) error {
	p.PanelName = ctx.Args().Get(0)
	if p.PanelName == "" {
		fmt.Println(errors.New("Error: Panel name is not set!"))
		return errors.New("Error: Panel name is not set!")

	}
	if p.DashboardName == "" {
		fmt.Println(errors.New("Error: Required flag \"dashboard name\"(-d) are not set!"))
		return errors.New("Error: Required flag \"dashboard name\"(-d) are not set!")
	}

	c, err := config.Read()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	req, _ := http.NewRequest("GET", c.Url+"/api/dashboards/db/"+p.DashboardName, nil)
	req.Header.Add("Authorization", "Bearer "+c.ApiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var dashboardExtended template.DashboardExtended
	err = json.Unmarshal([]byte(body), &dashboardExtended)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	p.printPanel(dashboardExtended)

	return nil
}

func (p *Panel) Flags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "d",
			Usage:       "specify dashboard name",
			Value:       "",
			Destination: &p.DashboardName,
		},
		cli.StringFlag{
			Name:        "r",
			Usage:       "specify row name",
			Value:       "",
			Destination: &p.RowName,
		},
	}

}

func (p *Panel) printPanel(dashboardExtended template.DashboardExtended) {
	counter := 0
	for _, row := range dashboardExtended.Dashboard.Rows {
		if p.RowName != "" && strings.ToLower(p.RowName) != strings.ToLower(row.Title) {
			continue
		}
		for _, panel := range row.Panels {
			if p.PanelName == panel["title"] {
				counter++
				if counter > 1 {
					fmt.Println("WARNING: multiple panel with same name found. Returning the first...")
					return
				}

				out, _ := yaml.Marshal(panel)
				fmt.Println(string([]byte(out)))
			}
		}
	}
}
