package get

import (
  "encoding/json"
  "errors"
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
  "strings"
  "text/tabwriter"

  "github.com/hajarbleh/grafcli/config"
  "github.com/hajarbleh/grafcli/template"
  "github.com/urfave/cli"
)

type Panels struct {
  DashboardName string
  RowName string
}

func (p *Panels) Execute(ctx *cli.Context) error {
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
  counter := 1
  for _, row := range dashboardExtended.Dashboard.Rows {
    if p.RowName != "" && strings.ToLower(p.RowName) != strings.ToLower(row.Title) {
      continue
    }
    for _, panel := range row.Panels {
      fmt.Fprintf(w, "%d\t%s\t%s\n", counter, panel["title"], row.Title)
      counter++
    }
  }
  fmt.Fprintln(w)
  w.Flush()
}
