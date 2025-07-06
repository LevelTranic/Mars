package gdel

import (
	"github.com/3JoB/unsafeConvert"
	"github.com/savsgio/atreugo/v11"

	"Mars/database/v2/controller"
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
		return helper.HandleError(c, "Project not found", 404)
	}

	version, err := controller.FindVersionByProjectAndName(project.ID, versionName)
	if err != nil {
		return helper.HandleError(c, "Version not found", 404)
	}

	if builder != "latest" {
		buildNumber = unsafeConvert.StringToInt(builder)
	} else {
		buildNumber = -1
	}
	if _, err := controller.FindBuildByProjectAndVersionAndNumber(project.ID, version.Name, buildNumber); err != nil {
		return helper.HandleError(c, "Build not found", 404)
	}

	if err := controller.DeleteBuild(project.ID, version.Name, buildNumber); err != nil {
		return helper.HandleError(c, "Deleting build failures: "+err.Error(), 500)
	}

	return c.JSONResponse(schemas.NewResult("Build Deleted"), 200)
}
