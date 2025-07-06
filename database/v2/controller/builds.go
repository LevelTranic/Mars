package controller

import (
	"Mars/database/shared"
	"Mars/database/v2/schemas"
)

func FindAllBuildsByProjectAndVersion(project, version string) ([]schemas.Build, error) {
	var builds []schemas.Build
	result := shared.DB.Where("project = ? AND version = ?", project, version).Find(&builds)
	return builds, result.Error
}

func FindAllBuildsByProjectAndVersionAndChannel(project, version, channel string) ([]schemas.Build, error) {
	var builds []schemas.Build
	result := shared.DB.Where("project = ? AND version = ? AND channel = ?", project, version, channel).Find(&builds)
	return builds, result.Error
}

func FindAllBuildsByProjectAndVersionIn(project string, versions []string) ([]schemas.Build, error) {
	var builds []schemas.Build
	result := shared.DB.Where("project = ? AND version IN ?", project, versions).Find(&builds)
	return builds, result.Error
}
