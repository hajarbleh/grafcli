package create

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

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

	c, err := config.Read()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	dashboard := template.NewDashboard(dName)
	dashboardExtended := template.DashboardExtended{Dashboard: dashboard}
	reqBody, err := json.Marshal(dashboardExtended)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(bytes.NewBuffer(reqBody))
	req, _ := http.NewRequest("POST", c.Url+"/api/dashboards/db", bytes.NewBuffer(reqBody))
	req.Header.Add("Authorization", "Bearer "+c.ApiKey)
	req.Header.Add("Content-Type", "Application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string([]byte(body)))

	return nil
}

func (d *Dashboard) Flags() []cli.Flag {
	return []cli.Flag{}
}
