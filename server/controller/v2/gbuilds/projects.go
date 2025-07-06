package gbuilds

import (
	"github.com/savsgio/atreugo/v11"

	"Mars/database/v2/controller"
	"Mars/server/helper"
	"Mars/server/schemas"
)

func Projects(c *atreugo.RequestCtx) error {
	defer func() {
		if p := recover(); p != nil {
			_ = helper.HandleInternalError(c, p)
			return
		}
	}()
	project := controller.FindAllProjects()
	proj := make([]string, 0, len(project))
	for _, v := range project {
		proj = append(proj, v.Name)
	}
	sche := schemas.ProjectsSchema{
		Projects: proj,
	}
	c.Response.Header.Set("Cache-Control", "public, max-age=604800")
	return c.JSONResponse(&sche, 200)
}
