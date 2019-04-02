package describe

import (
	//"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"gopkg.in/yaml.v2"
	//"github.com/ghodss/yaml"
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

	req, _ := http.NewRequest("GET", c.Url+"/api/dashboards/db/"+r.DashboardName, nil)
	req.Header.Add("Authorization", "Bearer "+c.ApiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode >= 300 {
		fmt.Println("Error: " + string([]byte(body)))
		return nil
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
