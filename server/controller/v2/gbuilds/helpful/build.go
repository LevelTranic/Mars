package helpful

import (
	dbsv2 "Mars/database/v2/schemas"
	"Mars/server/schemas"
)

func CreateBuildResponse(project dbsv2.Project, version dbsv2.Version, build dbsv2.Build) *schemas.ProjectVersionBuildSchema {
	changes := build.UnmarshalChanges()

	downloads := build.UnmarshalDownloads()

	return &schemas.ProjectVersionBuildSchema{
		ProjectRootSchema: schemas.ProjectRootSchema{
			ProjectId:   project.ID,
			ProjectName: project.Name,
		},
		Channel:   string(build.Channel),
		Version:   version.Name,
		Build:     build.Number,
		Time:      build.Time,
		Promoted:  build.Promoted,
		Changes:   changes,
		Downloads: downloads,
	}
}
