package gnew

import (
	"github.com/savsgio/atreugo/v11"

	"Mars/database/v2/controller"
	"Mars/server/helper"
	"Mars/server/schemas"
	"Mars/shared/utils/json"
)

func VersionFamily(c *atreugo.RequestCtx) error {
	defer func() {
		if p := recover(); p != nil {
			_ = helper.HandleInternalError(c, p)
			return
		}
	}()

	var jar *schemas.NewVersionGroupSchema
	_ = json.JSON.Unmarshal(c.Request.Body(), &jar)
	if jar == nil {
		return c.JSONResponse(schemas.NewError("missing parameter"), 400)
	}

	ok, info := controller.CreateFamily(jar.Project, jar.Name)
	if ok {
		return c.JSONResponse(schemas.NewResult(info), 200)
	}

	return c.JSONResponse(schemas.NewError(info), 400)
}
