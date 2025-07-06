package gdel

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"

	"github.com/3JoB/ulib/fsutil"
	"github.com/3JoB/unsafeConvert"
	"github.com/savsgio/atreugo/v11"

	"Mars/database/v2/controller"
	"Mars/server/helper"
	"Mars/server/schemas"
)

const (
	latestBuild    = "latest"
	latestBuildNum = -1
	storagePath    = "storage"
)

func Download(c *atreugo.RequestCtx) error {
	defer func() {
		if p := recover(); p != nil {
			_ = helper.HandleInternalError(c, p)
			return
		}
	}()

	projectName := c.UserValue("project").(string)
	versionName := c.UserValue("version").(string)
	buildID := c.UserValue("build").(string)
	downloadInfo := c.UserValue("downloadInfo").(string)

	project, err := controller.FindProjectByID(projectName)
	if err != nil {
		return helper.HandleError(c, "Project not found", 404)
	}

	version, err := controller.FindVersionByProjectAndName(project.ID, versionName)
	if err != nil {
		return helper.HandleError(c, "Version not found", 404)
	}

	buildNumber := latestBuildNum
	if buildID != latestBuild {
		buildNumber = unsafeConvert.StringToInt(buildID)
	}

	build, err := controller.FindBuildByProjectAndVersionAndNumber(project.ID, version.Name, buildNumber)
	if err != nil {
		return helper.HandleError(c, "Build not found", 404)
	}

	downloads := build.UnmarshalDownloads()
	if len(downloads) == 0 {
		return helper.HandleError(c, "Download not found", 404)
	}

	source, exists := downloads[downloadInfo]
	if !exists {
		return helper.HandleError(c, "Download not found", 404)
	}

	if source.Url == "" {
		if err := removeDownloadFile(project.Name, version.Name, build.Number, source.Name); err != nil {
			return helper.HandleError(c, "Unable to remove file: "+err.Error(), 500)
		}
	}

	delete(downloads, downloadInfo)
	if err := build.MarshalDownloads(downloads); err != nil {
		return helper.HandleError(c, err.Error(), 500)
	}

	if err := controller.CreateDownload(project.ID, version.Name, buildNumber, build); err != nil {
		return helper.HandleError(c, err.Error(), 500)
	}

	return c.JSONResponse(schemas.NewResult("Download Deleted"), 200)
}

func removeDownloadFile(projectName, versionName string, buildNum int, fileName string) error {
	filePath := filepath.Join(storagePath, projectName, versionName, strconv.Itoa(buildNum), fileName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return errors.New("download file not found")
	}
	return fsutil.Remove(filePath)
}
