package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Grafana defines Grafana HTTP client. Compatible w/ Grafana v4.5 API, usage w/
// later version may or may not work.
type Grafana struct {
	Host   string
	APIKey string
}

// NewGrafana creates new instance of Grafana client.
func NewGrafana(host, apiKey string) *Grafana {
	return &Grafana{host, apiKey}
}

// CreateDashboard is the create/update dashboard API. Param `dashboard` should
// be in JSON string format.
func (g *Grafana) CreateDashboard(dashboard string, overwrite bool, msg string) (respBody []byte, err error) {
	body := fmt.Sprintf(
		`{dashboard:%v,overwrite:%v,message,%v}`,
		dashboard, overwrite, msg,
	)
	return g.doHTTPReq(http.MethodPost, "/api/dashboards/db", []byte(body))
}

// GetDashboard is the get dashboard API.
func (g *Grafana) GetDashboard(slug string) (respBody []byte, err error) {
	return g.doHTTPReq(http.MethodGet, "/api/dashboards/db/"+slug, nil)
}

// SearchDashboards is the search dashboards API. Currently doesn't support any
// query params.
func (g *Grafana) SearchDashboards() (respBody []byte, err error) {
	return g.doHTTPReq(http.MethodGet, "/api/search?type=dash-db", nil)
}

func (g *Grafana) doHTTPReq(method, pth string, body []byte) (respBody []byte, err error) {
	req, _ := http.NewRequest(method, g.Host+pth, bytes.NewReader(body))
	req.Header.Add("Authorization", "Bearer "+g.APIKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	b, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("invalid response status %v: %v", resp.Status, string(b))
	}

	return b, nil
}
