package template

import()

type DashboardRow struct {
  Collapse bool `json:"collapse"`
  Height interface{} `json:"height"`
  //Panels []DashboardPanels `json:"panels"`
  Repeat *bool `json:"repeat"`
  RepeatIteration *int `json:"repeatIteration"`
  RepeatRowId *int `json:"repeatRowId"`
  ShowTitle bool `json:"showTitle"`
  Title string `json:"title"`
  TitleSize string `json:"titleSize"`
}
