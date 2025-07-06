package gnew

import (
	"github.com/savsgio/atreugo/v11"

	"Mars/database/v2/controller"
	"Mars/server/helper"
	"Mars/server/schemas"
	"Mars/shared/utils/json"
)

// Project route can be used to create new Projects
//
// @route /v2/new/project
//
// @params projectName: string
func Project(c *atreugo.RequestCtx) error {
	defer func() {
		if p := recover(); p != nil {
			_ = helper.HandleInternalError(c, p)
			return
		}
	}()
	var jar *schemas.NewProjectSchema
	_ = json.JSON.Unmarshal(c.Request.Body(), &jar)
	if jar == nil {
		return c.JSONResponse(schemas.NewError("missing parameter"), 400)
	}
	yes, info := controller.CreateProject(jar.ProjectId)
	if yes {
		return c.JSONResponse(schemas.NewResult("done"), 200)
	}
	return c.JSONResponse(schemas.NewErrors(info), 400)
}
