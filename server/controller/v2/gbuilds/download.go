package gbuilds

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/3JoB/ulib/litefmt"
	"github.com/3JoB/unsafeConvert"
	"github.com/gabriel-vasile/mimetype"
	"github.com/savsgio/atreugo/v11"

	"Mars/database/v2/controller"
	"Mars/server/helper"
	"Mars/server/schemas"
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
	buildNumber := 0
	downloadInfo := c.UserValue("downloadInfo")

	project, err := controller.FindProjectByID(projectName)
	if err != nil {
		return c.JSONResponse(schemas.NewError("Project not found"), 404)
	}

	version, err := controller.FindVersionByProjectAndName(project.ID, versionName)
	if err != nil {
		return c.JSONResponse(schemas.NewError("Version not found"), 404)
	}

	if buildID != "latest" {
		buildNumber = unsafeConvert.StringToInt(buildID)
	} else {
		buildNumber = -1
	}
	build, err := controller.FindBuildByProjectAndVersionAndNumber(project.ID, version.Name, buildNumber)
	if err != nil {
		return c.JSONResponse(schemas.NewError("Build not found"), 404)
	}

	downloads := build.UnmarshalDownloads()
	if len(downloads) == 0 {
		return c.JSONResponse(schemas.NewError("Download not found"), 404)
	}

	downloadSSL := ""
	if downloadInfo != nil {
		sDownloadInfo := downloadInfo.(string)
		dep, ok := downloads[sDownloadInfo]
		if ok {
			if dep.Url != "" {
				return c.RedirectResponse(dep.Url, 302)
			}
			downloadSSL = dep.Name
		} else {
			for _, d := range downloads {
				if d.Name == sDownloadInfo {
					if d.Url != "" {
						return c.RedirectResponse(d.Url, 302)
					}
					downloadSSL = d.Name
					break
				}
			}
			return c.JSONResponse(schemas.NewError("Download not found"), 404)
		}
	} else {
		for _, d := range downloads {
			if d.Url != "" {
				return c.RedirectResponse(d.Url, 302)
			}
			downloadSSL = d.Name
			break
		}
	}

	filePath := filepath.Join("storage", project.Name, version.Name, strconv.Itoa(build.Number), downloadSSL)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return c.JSONResponse(schemas.NewError("Download file not found"), 404)
	}

	c.Response.Header.Set("Content-Disposition", litefmt.Sprint("attachment; filename=", downloadSSL))
	c.Response.Header.Set("ETag", fmt.Sprintf("%x", time.Now().Unix()))

	mtype, _ := mimetype.DetectFile(filePath)

	return c.FileResponse(downloadSSL, filePath, mtype.String())
}
