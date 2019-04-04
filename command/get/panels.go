package get

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/hajarbleh/grafcli/client"
	"github.com/hajarbleh/grafcli/config"
	"github.com/hajarbleh/grafcli/template"
	"github.com/urfave/cli"
)

type Panels struct {
	DashboardName string
	RowName       string
}

func (p *Panels) Execute(ctx *cli.Context) error {
	if p.DashboardName == "" {
		fmt.Println(errors.New("Error: Required flag \"dashboard name\"(-d) are not set!"))
		return errors.New("Error: Required flag \"dashboard name\"(-d) are not set!")
	}

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

	p.printList(dashboardExtended)

	return nil
}

func (p *Panels) Flags() []cli.Flag {
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

func (p *Panels) printList(dashboardExtended template.DashboardExtended) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "NO.\tPANEL NAME\tROW NAME")
	counter := 0
	for _, row := range dashboardExtended.Dashboard.Rows {
		if p.RowName != "" && strings.ToLower(p.RowName) != strings.ToLower(row.Title) {
			continue
		}
		for _, panel := range row.Panels {
			counter++
			fmt.Fprintf(w, "%d\t%s\t%s\n", counter, panel["title"], row.Title)
		}
	}
	fmt.Fprintln(w)
	w.Flush()
}
