package gbuilds

import (
	"slices"

	"github.com/savsgio/atreugo/v11"

	"Mars/database/v2/controller"
	"Mars/server/controller/v2/gbuilds/helpful"
	"Mars/server/helper"
	"Mars/server/schemas"
)

func FamilyBuilds(c *atreugo.RequestCtx) error {
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
	versions := make([]string, len(version))
	builds := make([]schemas.BuildsSchema, len(version))
	if len(version) != 0 {
		for _, v := range version {
			versions = append(versions, v.Name)
			b, _ := controller.FindAllBuildsByProjectAndVersion(projectName, v.Name)
			if b == nil || len(b) == 0 {
				continue
			}
			bs := helpful.DatabaseBuildsToServerBuilds(b)
			builds = append(builds, bs...)
		}
		slices.Sort(versions)
	}

	r := schemas.FamilyVersionsBuildsSchema{
		FamilyVersionsSchema: schemas.FamilyVersionsSchema{
			ProjectRootSchema: schemas.ProjectRootSchema{
				ProjectName: projectName,
				ProjectId:   projectName,
			},
			VersionGroup: familyName,
			Versions:     versions,
		},
		Builds: builds,
	}

	return c.JSONResponse(r, 200)
}
