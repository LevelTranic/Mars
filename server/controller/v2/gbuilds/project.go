package gbuilds

import (
	"github.com/savsgio/atreugo/v11"
	"gorm.io/gorm"

	"Mars/database/v2/controller"
	"Mars/server/controller/v2/gbuilds/helpful"
	"Mars/server/helper"
	"Mars/server/schemas"
)

func Project(c *atreugo.RequestCtx) error {
	defer func() {
		if p := recover(); p != nil {
			_ = helper.HandleInternalError(c, p)
			return
		}
	}()
	proj := c.UserValue("project").(string)
	project, err := controller.FindProjectByID(proj)
	if err != nil {
		return c.JSONResponse(schemas.NewError("Project not found"), 404)
	}

	families, err := controller.FindAllFamilyByProject(project.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return c.JSONResponse(schemas.NewError("Failed to fetch version families"), 500)
	}

	versions, err := controller.FindAllVersionByProject(project.Name)
	if err != nil && err != gorm.ErrRecordNotFound {
		return c.JSONResponse(schemas.NewError("Failed to fetch versions"), 500)
	}

	response := helpful.CreateProjectResponse(project, families, versions)

	return c.JSONResponse(response, 200)
}
