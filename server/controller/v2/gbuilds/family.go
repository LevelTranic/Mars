package gbuilds

import (
	"slices"

	"github.com/savsgio/atreugo/v11"

	"Mars/database/v2/controller"
	"Mars/server/helper"
	"Mars/server/schemas"
)

func Family(c *atreugo.RequestCtx) error {
	defer func() {
		if p := recover(); p != nil {
			_ = helper.HandleInternalError(c, p)
			return
		}
	}()
	projectName := c.UserValue("project").(string)
	familyName := c.UserValue("family").(string)

	project, err := controller.FindProjectByID(projectName)
	if err != nil {
		return c.JSONResponse(schemas.NewError("Project not found"), 404)
	}

	version, err := controller.FindAllVersionByProjectAndGroup(project.ID, familyName)
	if err != nil {
		return c.JSONResponse(schemas.NewError("Version not found"), 404)
	}
	p := make([]string, len(version))
	if len(version) != 0 {
		for _, v := range version {
			p = append(p, v.Name)
		}
		slices.Sort(p)
	}
	r := schemas.FamilyVersionsSchema{
		ProjectRootSchema: schemas.ProjectRootSchema{
			ProjectId:   projectName,
			ProjectName: projectName,
		},
		VersionGroup: familyName,
		Versions:     p,
	}
	return c.JSONResponse(r, 200)
}
