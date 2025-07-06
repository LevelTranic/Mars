package controller

import (
	"Mars/database/shared"
	schemas2 "Mars/database/v2/schemas"
)

func CreateDownload(project, version string, build int, app schemas2.Build) error {
	/*, err := FindBuildByProjectAndVersionAndNumber(project, version, build)
	if err != nil {
		return err
	}*/
	return shared.DB.Where("project = ? AND version = ? AND number = ?", project, version, build).Updates(app).Error
}
