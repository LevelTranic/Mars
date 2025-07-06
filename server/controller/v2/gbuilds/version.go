package gbuilds

import (
	"github.com/grafana/regexp"
	"github.com/savsgio/atreugo/v11"
	"gorm.io/gorm"

	"Mars/database/v2/controller"
	"Mars/server/controller/v2/gbuilds/helpful"
	"Mars/server/helper"
	"Mars/server/schemas"
)

const VersionPattern = `^[a-zA-Z0-9.\-_]+$`

func Version(c *atreugo.RequestCtx) error {
	defer func() {
		if p := recover(); p != nil {
			_ = helper.HandleInternalError(c, p)
			return
		}
	}()
	projectName := c.UserValue("project").(string)
	versionName := c.UserValue("version").(string)

	// Validate input
	if !regexp.MustCompile("^[a-z]+$").MatchString(projectName) || !regexp.MustCompile(VersionPattern).MatchString(versionName) {
		return c.JSONResponse(schemas.NewError("Invalid input"), 400)
	}

	project, err := controller.FindProjectByID(projectName)
	if err != nil {
		return c.JSONResponse(schemas.NewError("Project not found"), 404)
	}

	version, err := controller.FindVersionByProjectAndName(project.ID, versionName)
	if err != nil {
		return c.JSONResponse(schemas.NewError("Version not found"), 404)
	}

	builds, err := controller.FindAllBuildsByProjectAndVersion(project.ID, version.Name)
	if err != nil && err != gorm.ErrRecordNotFound {
		return c.JSONResponse(schemas.NewError("Builds not found"), 404)
	}

	response := helpful.CreateVersionResponse(project, version, builds)

	// Cache the response for 5 minutes
	c.Response.Header.Set("Cache-Control", "public, max-age=300")
	return c.JSONResponse(response, 200)
}
