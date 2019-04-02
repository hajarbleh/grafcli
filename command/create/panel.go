package create

import (
  "bytes"
  "encoding/json"
  "errors"
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
  "strings"

  "github.com/hajarbleh/grafcli/config"
  "github.com/hajarbleh/grafcli/template"
  "github.com/hajarbleh/grafcli/utility"
  "github.com/urfave/cli"
)

type Panel struct {
  DashboardName string
  RowName string
  PanelName string
  PanelType string
}

func (p *Panel) Execute(ctx *cli.Context) error {
  if p.DashboardName == "" {
    fmt.Println(errors.New("Error: Required flag \"dashboard name\"(-d) are not set!"))
    return errors.New("Error: Required flag \"dashboard name\"(-d) are not set!")
  }

  if p.RowName == "" {
    fmt.Println(errors.New("Error: Required flag \"row name\"(-r) are not set!"))
    return errors.New("Error: Required flag \"row name\"(-r) are not set!")
  }

  c, err := config.Read()
  if err != nil {
    fmt.Println(err.Error())
    return err
  }

  interaction := utility.Interaction{
    Reader: os.Stdin,
  }

  p.PanelName, _ = interaction.AskUserInput("Enter panel name")
  p.PanelType, _ = interaction.AskUserInput("Enter panel type (graph, singlestat)")

  newPanel := make(map[string]interface{})

  switch p.PanelType {
  case "graph":
    newPanel = p.defaultGraphPanel()
  case "singlestat":
    // newPanel = p.defaultSinglestatPanel()
  default:
    fmt.Println("Error: unsupported panel type", p.PanelType)
    return nil
  }

  req, _ := http.NewRequest("GET", c.Url+"/api/dashboards/db/"+p.DashboardName, nil)
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
  newPanel["id"] = p.getMaxRowId(dashboardExtended.Dashboard.Rows)+1

  rowFound := 0
  for idx, r := range dashboardExtended.Dashboard.Rows {
    if strings.ToLower(r.Title) == strings.ToLower(p.RowName) {
      rowFound++
      if rowFound == 1 {
        dashboardExtended.Dashboard.Rows[idx].Panels = append(dashboardExtended.Dashboard.Rows[idx].Panels, newPanel)
      } else {
        fmt.Println("Warning: multiple row with name \"%s\" found. Executing only on first row found", p.RowName)
        break
      }
    }
  }

  if rowFound == 0 {
    fmt.Println("No row with name \"%s\" found. Cancelling create panel..", p.RowName)
    return nil
  }

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

  if resp.StatusCode >= 300 {
    fmt.Println("Error: " + string([]byte(body)))
    return nil
  }

  fmt.Println("Panel successfully created!")
  return nil
}

func (p *Panel) Flags() []cli.Flag {
  return []cli.Flag{
    cli.StringFlag{
      Name:        "d",
      Usage:       "specify dashboard name",
      Value:       "",
      Destination: &p.DashboardName,
    },
    cli.StringFlag{
      Name: "r",
      Usage: "specify row name",
      Value: "",
      Destination: &p.RowName,
    },
  }
}

func (p *Panel) defaultGraphPanel() map[string]interface{} {
  newPanel := make(map[string]interface{})
  newPanel["aliasColors"] = make(map[string]interface{})
  newPanel["bars"] = false
  newPanel["dashLength"] = 10
  newPanel["dashes"] = false
  newPanel["datasource"] = nil
  newPanel["fill"] = 1

  legend := make(map[string]interface{})
  legend["avg"] = false
  legend["current"] = false
  legend["max"] = false
  legend["min"] = false
  legend["show"] = true
  legend["total"] = false
  legend["values"] = false

  newPanel["legend"] = legend
  newPanel["lines"] = true
  newPanel["linewidth"] = 1
  newPanel["nullPointMode"] = "null"
  newPanel["percentage"] = false
  newPanel["pointradius"] = 5
  newPanel["points"] = false
  newPanel["renderer"] = "flot"
  newPanel["seriesOverrides"] = make([]interface{}, 0)
  newPanel["spaceLength"] = 10
  newPanel["span"] = 12
  newPanel["stack"] = false
  newPanel["steppedLine"] = false
  newPanel["targets"] = make([]map[string]interface{}, 1)
  newPanel["thresholds"] = make([]interface{}, 0)
  newPanel["timeFrom"] = nil
  newPanel["timeShift"] = nil
  newPanel["title"] = p.PanelName

  tooltip := make(map[string]interface{})
  tooltip["shared"] = true
  tooltip["sort"] = 0
  tooltip["value_type"] = "individual"

  newPanel["tooltip"] = tooltip
  newPanel["type"] = "graph"

  xaxis := make(map[string]interface{})
  xaxis["buckets"] = nil
  xaxis["mode"] = "time"
  xaxis["name"] = nil
  xaxis["show"] = true
  xaxis["values"] = make([]interface{}, 0)

  newPanel["xaxis"] = xaxis

  yaxes := make([]map[string]interface{}, 2)
  yaxes[0] = make(map[string]interface{})
  yaxes[0]["format"] = "short"
  yaxes[0]["label"] = nil
  yaxes[0]["logBase"] = 1
  yaxes[0]["max"] = nil
  yaxes[0]["min"] = nil
  yaxes[0]["show"] = true

  yaxes[1] = make(map[string]interface{})
  yaxes[1]["format"] = "short"
  yaxes[1]["label"] = nil
  yaxes[1]["logBase"] = 1
  yaxes[1]["max"] = nil
  yaxes[1]["min"] = nil
  yaxes[1]["show"] = true

  newPanel["yaxes"] = yaxes

  return newPanel
}

func (p *Panel) getMaxRowId(rows []template.DashboardRow) int {
  maxId := -1
  for _, r := range rows {
    for _, pp := range r.Panels {
      if maxId < int(pp["id"].(float64)) {
        maxId = int(pp["id"].(float64))
      }
    }
  }

  return maxId
}