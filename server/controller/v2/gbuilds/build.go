package gbuilds

import (
	"github.com/3JoB/unsafeConvert"
	"github.com/savsgio/atreugo/v11"

	"Mars/database/v2/controller"
	"Mars/server/controller/v2/gbuilds/helpful"
	"Mars/server/helper"
	"Mars/server/schemas"
)

func Build(c *atreugo.RequestCtx) error {
	defer func() {
		if p := recover(); p != nil {
			_ = helper.HandleInternalError(c, p)
			return
		}
	}()

	projectName := c.UserValue("project").(string)
	versionName := c.UserValue("version").(string)
	builder := c.UserValue("build").(string)
	buildNumber := 0

	project, err := controller.FindProjectByID(projectName)
	if err != nil {
		return c.JSONResponse(schemas.NewError("Project not found"), 404)
	}

	version, err := controller.FindVersionByProjectAndName(project.ID, versionName)
	if err != nil {
		return c.JSONResponse(schemas.NewError("Version not found"), 404)
	}

	if builder != "latest" {
		buildNumber = unsafeConvert.StringToInt(builder)
	} else {
		buildNumber = -1
	}
	build, err := controller.FindBuildByProjectAndVersionAndNumber(project.ID, version.Name, buildNumber)
	if err != nil {
		return c.JSONResponse(schemas.NewError("Build not found"), 404)
	}

	response := helpful.CreateBuildResponse(project, version, build)

	c.Response.Header.Set("Cache-Control", "public, max-age=604800")
	return c.JSONResponse(response, 200)
}
