package save

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ghodss/yaml"
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

	body, _ := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

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

	for idx, d := range dashboardExtended.Dashboard.Rows {
		if strings.ToLower(d.Title) == strings.ToLower(newRow.Title) {
			dashboardExtended.Dashboard.Rows[idx] = newRow
			break
		}
	}

	jsonBody, _ := json.Marshal(dashboardExtended)
	req, _ = http.NewRequest("POST", c.Url+"/api/dashboards/db", bytes.NewBuffer([]byte(jsonBody)))
	req.Header.Add("Authorization", "Bearer "+c.ApiKey)
	req.Header.Add("Content-Type", "application/json")
	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer resp.Body.Close()

	body, _ = ioutil.ReadAll(resp.Body)
	if resp.StatusCode >= 300 {
		fmt.Println("Error: " + string([]byte(body)))
		return nil
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
