package template

import(
  "time"
)

type DashboardMeta struct {
  Type string `json:"type"`
  CanSave bool `json:"canSave"`
  CanEdit bool `json:"canEdit"`
  CanStar bool `json:"canStar"`
  Slug string `json:"slug"`
  Expires time.Time `json:"expires"`
  Created time.Time `json:"created"`
  Updated time.Time `json:"updated"`
  UpdatedBy string `json:"updatedBy"`
  CreatedBy string `json:"createdBy"`
  Version int `json:version`
}