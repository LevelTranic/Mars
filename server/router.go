package server

import (
	"github.com/savsgio/atreugo/v11"

	"Mars/server/controller"
	_default "Mars/server/controller/default"
	"Mars/server/controller/v2/gbuilds"
	"Mars/server/controller/v2/gdel"
	"Mars/server/controller/v2/gnew"
)

func Router(s *atreugo.Atreugo) {
	s.ANY("/", _default.Index)
	v2 := s.NewGroupPath("/v2")

	nw := v2.NewGroupPath("/new").UseBefore(controller.Auth)
	nw.POST("/project", gnew.Project)
	nw.POST("/version", gnew.Version)
	nw.POST("/version_family", gnew.VersionFamily)
	nw.POST("/build", gnew.Build)
	nw.POST("/download", gnew.Download)
	nw.POST("/external_download", gnew.ExternalDownload)

	v2.GET("/projects", gbuilds.Projects)
	v2.GET("/projects/{project}", gbuilds.Project)
	v2.GET("/projects/{project}/version_group/{family}", gbuilds.Family)
	v2.GET("/projects/{project}/version_group/{family}/builds", gbuilds.FamilyBuilds)
	v2.GET("/projects/{project}/versions/{version}", gbuilds.Version)
	v2.GET("/projects/{project}/versions/{version}/builds", gbuilds.Builds)
	v2.GET("/projects/{project}/versions/{version}/builds/{build}", gbuilds.Build)
	v2.GET("/projects/{project}/versions/{version}/builds/{build}/downloads", gbuilds.Download)
	v2.GET("/projects/{project}/versions/{version}/builds/{build}/downloads/{downloadInfo}", gbuilds.Download)

	deleteRoute := v2.NewGroupPath("/projects").UseBefore(controller.Auth)
	deleteRoute.DELETE("/{project}/versions/{version}/builds/{build}", gdel.Build)
	deleteRoute.DELETE("/{project}/versions/{version}/builds/{build}/downloads/{downloadInfo}", gdel.Download)
}
