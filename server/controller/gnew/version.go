package gnew

import (
	"github.com/savsgio/atreugo/v11"

	"Mars/database/controller"
	"Mars/server/helper"
	"Mars/server/schemas"
	"Mars/shared/utils/json"
)

func Version(c *atreugo.RequestCtx) error {
	defer func() {
		if p := recover(); p != nil {
			_ = helper.HandleInternalError(c, p)
			return
		}
	}()

	var jar *schemas.NewVersionSchema
	_ = json.JSON.Unmarshal(c.Request.Body(), &jar)
	if jar == nil {
		return c.JSONResponse(schemas.NewError("missing parameter"), 400)
	}

	if err := controller.CreateVersion(jar.Project, jar.Version, jar.Group); err == nil {
		return c.JSONResponse(schemas.NewResult("done"), 200)
	} else {
		return c.JSONResponse(schemas.NewErrors(err), 400)
	}
}
