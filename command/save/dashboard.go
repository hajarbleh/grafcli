package save

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hajarbleh/grafcli/config"
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

	data, err := ioutil.ReadFile(d.Filename)
	if err != nil {
		return errors.New(fmt.Sprintf("Fatal error config file: %s \n", err))
	}
	jsonBody := bytes.NewBuffer([]byte(data))

	c, err := config.Read()
	if err != nil {
		return err
	}

	req, _ := http.NewRequest("POST", c.Url+"/api/dashboards/db", jsonBody)
	req.Header.Add("Authorization", "Bearer "+c.ApiKey)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Indent(buf, body, "", "    ")
	if err != nil {
		return err
	}

	return nil
}

func (d *Dashboard) Flags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "f",
			Usage:       "filename to save the resource",
			Value:       "",
			Destination: &d.Filename,
		},
	}
}
