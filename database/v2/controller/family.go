package controller

import (
	"time"

	"gorm.io/gorm"

	"Mars/database/shared"
	"Mars/database/v2/schemas"
)

func newFamily(project, name string) *gorm.DB {
	now := time.Now().UTC()
	versionFamily := schemas.VersionFamily{
		Project: project,
		Name:    name,
		Time:    &now,
	}
	return shared.DB.Create(&versionFamily)
}

func CreateFamily(project, name string) (bool, string) {
	var versionFamily schemas.VersionFamily
	if err := shared.DB.Where("project = ? AND name = ?", project, name).First(&versionFamily).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if err = newFamily(project, name).Error; err != nil {
				return false, err.Error()
			}
			return true, "done"
		}
		return false, err.Error()
	}
	return false, "The same content already exists"
}

// FindFamilyByProjectAndName finds a VersionFamily by project and name
func FindFamilyByProjectAndName(project, name string) (schemas.VersionFamily, error) {
	var versionFamily schemas.VersionFamily
	result := shared.DB.Where("project = ? AND name = ?", project, name).First(&versionFamily)
	return versionFamily, result.Error
}

// FindAllFamilyByProject finds all VersionFamilies by project
func FindAllFamilyByProject(project string) ([]schemas.VersionFamily, error) {
	var versionFamilies []schemas.VersionFamily
	result := shared.DB.Where("project = ?", project).Find(&versionFamilies)
	return versionFamilies, result.Error
}
