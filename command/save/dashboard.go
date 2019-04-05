package save

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/hajarbleh/grafcli/client"
	"github.com/hajarbleh/grafcli/config"
	"github.com/hajarbleh/grafcli/utility"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

type Dashboard struct {
	Filename string
}

func (d *Dashboard) Execute(ctx *cli.Context) error {
	if d.Filename == "" {
		fmt.Println(errors.New("Error: must specify file name!"))
		return errors.New("must specify file name")
	}

	data, err := d.readDashboard(d.Filename)
	if err != nil {
		fmt.Println(err)
		return err
	}

	jsonDashboard, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	grafana := client.NewGrafana(config.URL, config.APIKey)
	if _, err := grafana.CreateDashboard(string(jsonDashboard), false, "Updated by grafcli"); err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Dashboard successfully saved!")
	return nil
}

func (d *Dashboard) readDashboard(pth string) (map[string]interface{}, error) {
	data, err := d.readYAMLtoMap(pth)
	if err != nil {
		return nil, errors.Wrap(err, "error reading dashboard file")
	}

	dir := strings.TrimSuffix(pth, filepath.Ext(pth))
	rowPths, err := utility.ListFiles(dir)
	if err != nil {
		fmt.Println("Rows for dashboard not found.")
	}

	rows := make([]interface{}, len(rowPths))
	for i, rowPth := range rowPths {
		row, err := d.readRow(rowPth)
		if err != nil {
			return nil, errors.Wrap(err, "error reading row "+rowPth)
		}
		rows[i] = row
	}
	data["rows"] = rows
	return data, nil
}

func (d *Dashboard) readRow(pth string) (map[string]interface{}, error) {
	data, err := d.readYAMLtoMap(pth)
	if err != nil {
		return nil, errors.Wrap(err, "error reading row file")
	}

	dir := strings.TrimSuffix(pth, filepath.Ext(pth))
	panelPths, err := utility.ListFiles(dir)
	if err != nil {
		fmt.Printf("Panels for row '%s' not found.\n", filepath.Base(pth))
	}

	panels := make([]interface{}, len(panelPths))
	for i, panelPth := range panelPths {
		panel, err := d.readPanel(panelPth)
		if err != nil {
			return nil, errors.Wrap(err, "error reading panel "+panelPth)
		}
		panels[i] = panel
	}
	data["panels"] = panels
	return data, nil
}

func (d *Dashboard) readPanel(pth string) (map[string]interface{}, error) {
	data, err := d.readYAMLtoMap(pth)
	if err != nil {
		return nil, errors.Wrap(err, "error reading panel file")
	}
	return data, nil
}

func (d *Dashboard) readYAMLtoMap(pth string) (map[string]interface{}, error) {
	b, err := ioutil.ReadFile(pth)
	if err != nil {
		return nil, err
	}

	data := map[string]interface{}{}
	if err := yaml.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return data, err
}

func (d *Dashboard) Flags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "f",
			Usage:       "Dashboard YAML filename to be saved to Grafana",
			Value:       "",
			Destination: &d.Filename,
		},
	}
}
