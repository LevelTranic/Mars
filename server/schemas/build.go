package schemas

import (
	"time"

	dbsv2 "Mars/database/v2/schemas"
	hssv2 "Mars/shared/httpschemas/v2"
)

type NewBuildSchema struct {
	Project  string         `json:"project"`
	Version  string         `json:"version"`
	Family   string         `json:"family"`
	Channel  dbsv2.Channel  `json:"channel"`
	Promoted bool           `json:"promoted"`
	Changes  []dbsv2.Change `json:"changes,omitempty"`
}

type NewBuildResult struct {
	Project string `json:"project"`
	Version string `json:"version"`
	Family  string `json:"family"`
	BuildID int    `json:"build_id"`
}

type BuildsSchema struct {
	Build     int                                        `json:"build"`
	Time      time.Time                                  `json:"time"`
	Channel   string                                     `json:"channel"`
	Promoted  bool                                       `json:"promoted"`
	Changes   []dbsv2.Change                             `json:"changes"`
	Downloads map[string]hssv2.ApplicationVersionsSchema `json:"downloads"`
}

type ProjectVersionBuildSchema struct {
	ProjectRootSchema
	Channel   string                                     `json:"channel"`
	Version   string                                     `json:"version"`
	Build     int                                        `json:"build"`
	Promoted  bool                                       `json:"promoted"`
	Time      time.Time                                  `json:"time"`
	Changes   []dbsv2.Change                             `json:"changes"`
	Downloads map[string]hssv2.ApplicationVersionsSchema `json:"downloads"`
}
