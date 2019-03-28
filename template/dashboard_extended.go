package template

import ()

type DashboardExtended struct {
	Dashboard Dashboard `json:"dashboard"`
	Overwrite bool      `json:"overwrite"`
	Message   string    `json:"message"`
}
