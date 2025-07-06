package v2

import (
	"runtime"

	"github.com/3JoB/ulib/litefmt"
)

const (
	CoreVersion       = "v1.4.3"
	PrivateAPIVersion = "v1.4.3"
	PublicAPIVersion  = "v1.4.2"
	DatabaseVersion   = "v1.3.2"
)

type Index struct {
	Server   string       `json:"server"`
	Github   string       `json:"github"`
	Docs     string       `json:"docs"`
	Platform string       `json:"platform"`
	Version  IndexVersion `json:"version"`
}

type IndexVersion struct {
	Core       string `json:"core"`
	PrivateAPI string `json:"private_api"`
	PublicAPI  string `json:"public_api"`
	Database   string `json:"database"`
}

var DefaultIndex = Index{
	Server:   "Tranic Mars",
	Github:   "https://github.com/404Setup/Mars",
	Docs:     "https://github.com/404Setup/Mars/wiki/Mars-API",
	Platform: litefmt.Sprint(runtime.GOOS, "/", runtime.GOARCH),
	Version: IndexVersion{
		Core:       CoreVersion,
		PrivateAPI: PrivateAPIVersion,
		PublicAPI:  PublicAPIVersion,
		Database:   DatabaseVersion,
	},
}
