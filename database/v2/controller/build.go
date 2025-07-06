package controller

import (
	"errors"
	"log/slog"
	"time"

	"gorm.io/gorm"

	"Mars/database/shared"
	"Mars/database/v2/schemas"
	schemas2 "Mars/server/schemas"
)

func DeleteBuild(project, version string, number int) error {
	return shared.DB.Where("project = ? AND version = ? AND number = ?", project, version, number).Delete(&schemas.Build{}).Error
}

func CreateBuild(s *schemas2.NewBuildSchema) (int, error) {
	var build schemas.Build
	var buildNum int

	// Calculate BuildNum here.
	// Take the buildNum of the last piece of data as a reference,
	// and add one to it as the buildNum of the new data created.
	result := shared.DB.Where("project = ? AND version = ?", s.Project, s.Version).Last(&build)
	if result.Error == gorm.ErrRecordNotFound {
		buildNum = 1
	} else if result.Error != nil {
		return 0, errors.Join(shared.ProjectAndVersionNotFoundError, result.Error)
	} else {
		buildNum = build.Number + 1
	}

	times := time.Now().UTC()

	newBuild := &schemas.Build{
		Project:   s.Project,
		Version:   s.Version,
		Number:    buildNum,
		Time:      times,
		Changes:   nil,
		Downloads: nil,
		Channel:   s.Channel,
		Promoted:  s.Promoted,
	}

	if s.Changes != nil || len(s.Changes) != 0 {
		_ = newBuild.MarshalChanges(s.Changes)
	}

	return buildNum, shared.DB.Create(newBuild).Error
}

func FindBuildByProjectAndVersionAndNumber(project, version string, number int) (schemas.Build, error) {
	var build schemas.Build
	var result *gorm.DB
	if number == -1 {
		result = shared.DB.Where("project = ? AND version = ?", project, version).Last(&build)
	} else {
		result = shared.DB.Where("project = ? AND version = ? AND number = ?", project, version, number).First(&build)
	}
	return build, result.Error
}

func RemoveAllBundlerBuilds() {
	projs := FindAllProjects()
	for _, proj := range projs {
		versions, err := FindAllVersionByProject(proj.Name)
		if err != nil {
			slog.Error(err.Error())
			continue
		}
		for _, version := range versions {
			builds, err := FindAllBuildsByProjectAndVersion(proj.Name, version.Name)
			if err != nil {
				slog.Error(err.Error())
				continue
			}
			for _, build := range builds {
				d := build.UnmarshalDownloads()
				_, ok := d["bundler"]
				if !ok {
					continue
				}
				delete(d, "bundler")
				_ = build.MarshalDownloads(d)
				shared.DB.Where("project = ? AND version = ?", proj.Name, version.Name).Updates(&build)
			}
		}
	}
}
