package gbuilds

import (
	"github.com/3JoB/unsafeConvert"
	"github.com/savsgio/atreugo/v11"
	"gorm.io/gorm"

	"Mars/database/v2/controller"
	dbsv2 "Mars/database/v2/schemas"
	"Mars/server/controller/v2/gbuilds/helpful"
	"Mars/server/helper"
	"Mars/server/schemas"
)

func Builds(c *atreugo.RequestCtx) error {
	defer func() {
		if p := recover(); p != nil {
			_ = helper.HandleInternalError(c, p)
			return
		}
	}()
	projectName := c.UserValue("project").(string)
	versionName := c.UserValue("version").(string)

	project, err := controller.FindProjectByID(projectName)
	if err != nil {
		return c.JSONResponse(schemas.NewError("Project not found"), 404)
	}

	version, err := controller.FindVersionByProjectAndName(project.ID, versionName)
	if err != nil {
		return c.JSONResponse(schemas.NewError("Version not found"), 404)
	}

	channel := c.QueryArgs().Peek("channel")
	var builds []dbsv2.Build
	if channel != nil {
		builds, err = controller.FindAllBuildsByProjectAndVersionAndChannel(
			project.ID,
			version.Name,
			unsafeConvert.StringPointer(channel))
	} else {
		builds, err = controller.FindAllBuildsByProjectAndVersion(project.ID, version.Name)
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return c.JSONResponse(schemas.NewError("Builds not found"), 404)
	}

	response := helpful.CreateBuildsResponse(project, version, builds)

	// Cache the response for 5 minutes
	c.Response.Header.Set("Cache-Control", "public, max-age=300")
	return c.JSONResponse(response, 200)
}
