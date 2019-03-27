package get

import (
  "encoding/json"
  "errors"
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
  "text/tabwriter"

  "github.com/hajarbleh/grafcli/config"
  "github.com/hajarbleh/grafcli/template"
  "github.com/urfave/cli"
)

type Rows struct {
  DashboardName string
}

func(r *Rows) Execute(ctx *cli.Context) error {
  if r.DashboardName == "" {
    fmt.Println("must specify dashboard name")
    return errors.New("must specify dashboard name")
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

  var dashboardExtended template.DashboardExtended
  err = json.Unmarshal([]byte(body), &dashboardExtended)
  if err != nil {
    fmt.Println(err.Error())
    return err
  }

  printList(dashboardExtended.Dashboard.Rows)

  return nil
}

func(r *Rows) Flags() []cli.Flag {
  return []cli.Flag{
    cli.StringFlag{
      Name: "d",
      Usage: "specify dashboard name",
      Value: "",
      Destination: &r.DashboardName,
    },
  }

}

func printList(dashboardRow []template.DashboardRow) {
  w := new(tabwriter.Writer)
  w.Init(os.Stdout, 0, 8, 0, '\t', 0)
  fmt.Fprintln(w, "NO.\tTITLE")
  for idx, d := range dashboardRow {
    fmt.Fprintf(w, "%d\t%s\n", idx+1, d.Title)
  }
  fmt.Fprintln(w)
  w.Flush()
}
