package template

import ()

type Dashboard struct {
  Id          int      `json:"id"`
  Title       string   `json:"title"`
  Tags        []string `json:"tags"`
  Timezone    string `json:"timezone"`
  // Rows        []Rows
  SchemaVersion int `json:"schemaVersion"`
  Version     int `json:"version"`

}

func NewDashboard(title string) Dashboard {
  dashboard := Dashboard{}
  dashboard.Title = title
  dashboard.Timezone = "browser"
  return dashboard
}