package controller

import (
	"Mars/database/shared"
	"Mars/database/v2/schemas"
)

// FindAllVersionByProject finds all Versions by project
func FindAllVersionByProject(project string) ([]schemas.Version, error) {
	var versions []schemas.Version
	result := shared.DB.Where("project = ?", project).Find(&versions)
	return versions, result.Error
}

// FindAllVersionByProjectID finds all Versions by project
func FindAllVersionByProjectID(projectId string) ([]schemas.Version, error) {
	var versions []schemas.Version
	result := shared.DB.Where("name = ?", projectId).Find(&versions)
	return versions, result.Error
}

// FindAllVersionByProjectAndGroup finds all Versions by project and group
func FindAllVersionByProjectAndGroup(project, group string) ([]schemas.Version, error) {
	var versions []schemas.Version
	result := shared.DB.Where("project = ? AND `group` = ?", project, group).Find(&versions)
	return versions, result.Error
}
