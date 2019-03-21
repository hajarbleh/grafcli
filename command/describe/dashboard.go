package describe

import (
  "bytes"
  "encoding/json"
  "errors"
  "fmt"
  "io/ioutil"
  "net/http"

  "github.com/ghodss/yaml"
  "github.com/hajarbleh/grafcli/config"
  "github.com/urfave/cli"
)

type Dashboard struct {
}

func (d *Dashboard) Execute(ctx *cli.Context) error {
  dName := ctx.Args().Get(0)
  if dName == "" {
    fmt.Println("must specify dashboard name")
    return errors.New("must specify dashboard name")
  }

  c, err := config.Read()
  if err != nil {
    return err
  }

  req, _ := http.NewRequest("GET", c.Url + "/api/dashboards/db/" + dName, nil)
  req.Header.Add("Authorization", "Bearer " + c.ApiKey)
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    fmt.Println(err.Error())
    return err
  }

  defer resp.Body.Close()

  buf := new(bytes.Buffer)
  body, _ := ioutil.ReadAll(resp.Body)


  err = json.Indent(buf, body, "", "    ")
  y, err := yaml.JSONToYAML([]byte(body))
  if err != nil {
    fmt.Println(err.Error())
    return err
  }

  fmt.Println(string([]byte(y)))

  return nil
}

func (d *Dashboard) Flags() []cli.Flag {
  return []cli.Flag{
  }
}
