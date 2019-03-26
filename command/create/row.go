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

  body, _ := ioutil.ReadAll(resp.Body)

  resp.Body.Close()

  if resp.StatusCode >= 300 {
    fmt.Println("Error: "+string([]byte(body)))
    return nil
  }

  var dashboardExtended template.DashboardExtended
  err = json.Unmarshal([]byte(body), &dashboardExtended)

  if err != nil {
    fmt.Println(err.Error())
    return err
  }

  newRow := template.NewDashboardRow(rName)
  dashboardExtended.Dashboard.Rows = append(dashboardExtended.Dashboard.Rows, newRow)

  jsonBody, err := json.Marshal(dashboardExtended)
  if err != nil {
    fmt.Println(err.Error())
    return err
  }

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

  buf := new(bytes.Buffer)
  body, _ = ioutil.ReadAll(resp.Body)
  err = json.Indent(buf, body, "", "    ")
  if err != nil {
    fmt.Println(err)
    return err
  }

  fmt.Println("Dashboard successfully saved!")

  return nil
}

func (r *Row) Flags() []cli.Flag {
  return []cli.Flag{
    cli.StringFlag{
      Name: "d",
      Usage: "specify dashboard name",
      Value: "",
      Destination: &r.DashboardName,
    },
  }
}
