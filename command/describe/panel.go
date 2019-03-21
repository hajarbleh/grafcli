package describe

import (
  "errors"
  "fmt"
  "io/ioutil"
  "net/http"

  "github.com/hajarbleh/grafcli/config"
  "github.com/urfave/cli"
)

type Panel struct {
  DashboardName string
  PanelName string
}

func (p *Panel) Execute(ctx *cli.Context) error {
  if p.PanelName == "" {
    return errors.New("must specify panel name!")
  }
  if p.DashboardName == "" {
    return errors.New("must specify dashboard name!")
  }

  c, err := config.Read()
  if err != nil {
    return err
  }

  req, _ := http.NewRequest("GET", c.Url+"/api/search", nil)
  q := req.URL.Query()
  q.Add("query", p.DashboardName)
  q.Add("type", "dash-db")
  req.URL.RawQuery = q.Encode()
  fmt.Println(req.URL.String())
  req.Header.Add("Authorization", "Bearer "+c.ApiKey)
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return err
  }

  defer resp.Body.Close()

  body, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string([]byte(body)))
  return nil
}

func (p *Panel) Flags() []cli.Flag {
  return []cli.Flag{
    cli.StringFlag{
      Name: "d",
      Usage: "specify dashboard name",
      Value: "",
      Destination: &p.DashboardName,
    },
    cli.StringFlag{
      Name: "p",
      Usage: "specify panel name",
      Value: "",
      Destination: &p.PanelName,
    },
  }
}
