package helpful

import (
	dbsv2 "Mars/database/v2/schemas"
	"Mars/server/schemas"
)

func CreateVersionResponse(project dbsv2.Project, version dbsv2.Version, builds []dbsv2.Build) *schemas.ProjectVersionsSchema {
	buildNumbers := make([]int, len(builds))
	for i, build := range builds {
		buildNumbers[i] = build.Number
	}
	return &schemas.ProjectVersionsSchema{
		ProjectRootSchema: schemas.ProjectRootSchema{
			ProjectId:   project.ID,
			ProjectName: project.Name,
		},
		Version: version.Name,
		Builds:  buildNumbers,
	}
}
