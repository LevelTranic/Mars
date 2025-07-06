package v2

type NewDownloadSchema struct {
	Project string                               `json:"project"`
	Version string                               `json:"version"`
	Build   int                                  `json:"build"`
	File    map[string]ApplicationVersionsSchema `json:"file,omitempty"`
}

type ApplicationVersionsSchema struct {
	Name   string `json:"name,omitempty"`
	Sha256 string `json:"sha256,omitempty"`
	Url    string `json:"url,omitempty"`
}

var ApplicationVersionsNil = ApplicationVersionsSchema{}
