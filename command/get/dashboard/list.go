package dashboard

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/hajarbleh/grafcli/client"
	"github.com/hajarbleh/grafcli/config"
	"github.com/hajarbleh/grafcli/template"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

type List struct {
}

func (l *List) Execute(ctx *cli.Context) error {
	c, err := config.Read()
	if err != nil {
		fmt.Println("Error loading configuration")
		return errors.Wrap(err, "Error loading configuration")
	}

	grafana := client.NewGrafana(c.Url, c.ApiKey)
	body, err := grafana.SearchDashboards()
	if err != nil {
		fmt.Println(err)
		return err
	}

	var dashboardList []template.DashboardList
	err = json.Unmarshal([]byte(body), &dashboardList)
	err = printList(dashboardList)

	return nil
}

func printList(dList []template.DashboardList) error {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "ID\tUID\tTITLE\tURL\tTYPE\tTAGS")
	for _, d := range dList {
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\t%s\t\n", d.Id, d.Uid, d.Title, d.Url, d.Type, strings.Join(d.Tags, ","))
	}
	fmt.Fprintln(w)
	w.Flush()

	return nil
}
