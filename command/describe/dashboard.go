package describe

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/ghodss/yaml"
	"github.com/hajarbleh/grafcli/client"
	"github.com/hajarbleh/grafcli/config"
	"github.com/hajarbleh/grafcli/utility"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

type Dashboard struct {
	OutputDir string
}

func (d *Dashboard) Execute(ctx *cli.Context) error {
	dName := ctx.Args().Get(0)
	if dName == "" {
		fmt.Println("must specify dashboard name")
		return errors.New("must specify dashboard name")
	}

	grafana := client.NewGrafana(config.URL, config.APIKey)
	body, err := grafana.GetDashboard(dName)
	if err != nil {
		fmt.Println(err)
		return err
	}

	data := map[string]interface{}{}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		return err
	}

	dashboard := data["dashboard"].(map[string]interface{})
	name := utility.SanitizeFilename(dashboard["title"].(string))
	if err := d.writeDashboard(dashboard, filepath.Join(d.OutputDir, name)); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (d *Dashboard) writeDashboard(data map[string]interface{}, outPth string) error {
	rows := data["rows"].([]interface{})
	delete(data, "rows") // exclude row data from dashboard YAML.
	if err := d.writeMapToYAML(data, outPth+".yml"); err != nil {
		return errors.Wrap(err, "error saving dashboard data")
	}

	for i, rowi := range rows {
		row := rowi.(map[string]interface{})
		name := d.outputFilename(row["title"].(string), i, len(rows))
		if err := d.writeRow(row, filepath.Join(outPth, name)); err != nil {
			return err
		}
	}
	return nil
}

func (d *Dashboard) writeRow(data map[string]interface{}, outPth string) error {
	panels := data["panels"].([]interface{})
	delete(data, "panels") // exclude panel data from row YAML.
	if err := d.writeMapToYAML(data, outPth+".yml"); err != nil {
		return errors.Wrap(err, "error saving row data")
	}

	for i, paneli := range panels {
		panel := paneli.(map[string]interface{})
		name := d.outputFilename(panel["title"].(string), i, len(panels))
		if err := d.writePanel(panel, filepath.Join(outPth, name)); err != nil {
			return err
		}
	}
	return nil
}

func (d *Dashboard) writePanel(data map[string]interface{}, outPth string) error {
	return d.writeMapToYAML(data, outPth+".yml")
}

func (d *Dashboard) writeMapToYAML(data map[string]interface{}, outPth string) error {
	yml, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	return utility.WriteFile(yml, outPth)
}

// outputFilename generates proper filename for output YAMLs.
func (d *Dashboard) outputFilename(title string, order, totalItem int) string {
	title = utility.SanitizeFilename(title)

	// add number prefix to retain actual ordering in Grafana.
	pad := fmt.Sprintf("%v", len(fmt.Sprintf("%v", totalItem)))
	return fmt.Sprintf("%0"+pad+".f_%s", float64(order), title)
}

func (d *Dashboard) Flags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "o",
			Usage:       "Output directory",
			Value:       ".",
			Destination: &d.OutputDir,
		},
	}
}
