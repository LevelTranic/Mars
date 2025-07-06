package gnew

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/3JoB/ulib/fsutil"
	fshash "github.com/3JoB/ulib/fsutil/hash"
	"github.com/3JoB/unsafeConvert"
	"github.com/savsgio/atreugo/v11"
	"gorm.io/gorm"

	"Mars/database/v2/controller"
	"Mars/server/helper"
	"Mars/server/schemas"
	hssv2 "Mars/shared/httpschemas/v2"
	"Mars/shared/utils"
)

func Download(c *atreugo.RequestCtx) error {
	defer func() {
		if p := recover(); p != nil {
			_ = helper.HandleInternalError(c, p)
			return
		}
		utils.GC()
	}()

	project, version, build := "", "", ""
	{
		p1 := c.QueryArgs().Peek("project")
		p2 := c.QueryArgs().Peek("version")
		p3 := c.QueryArgs().Peek("build")
		if p1 == nil || p2 == nil || p3 == nil {
			return c.JSONResponse(schemas.NewError("project or version or build is nil"), 400)
		}

		project = unsafeConvert.StringPointer(p1)
		version = unsafeConvert.StringPointer(p2)
		build = unsafeConvert.StringPointer(p3)

		if unsafeConvert.StringToInt(build) == 0 {
			return c.JSONResponse(schemas.NewError("buildID can no longer be 0"), 400)
		}
	}

	{
		fp1 := filepath.Join("storage", project)
		if !fsutil.IsExist(fp1) {
			fsutil.Mkdir(fp1)
		}
		fp2 := filepath.Join(fp1, version)
		if !fsutil.IsExist(fp2) {
			fsutil.Mkdir(fp2)
		}
		fp3 := filepath.Join(fp2, build)
		if !fsutil.IsExist(fp3) {
			fsutil.Mkdir(fp3)
		}
	}

	builds, err := controller.FindBuildByProjectAndVersionAndNumber(project, version, unsafeConvert.StringToInt(build))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSONResponse(schemas.NewError("project or version or buildId is not found"), 400)
		}
		return c.JSONResponse(schemas.NewErrors(err), 500)
	}

	downloads := builds.UnmarshalDownloads()

	mf, err := c.MultipartForm()
	if err != nil {
		return c.JSONResponse(schemas.NewErrors(err), 500)
	}
	if mf.File == nil {
		return c.JSONResponse(schemas.NewError("file not found"), 500)
	}
	defer mf.RemoveAll()
	basicPath := filepath.Join("storage", project, version, build)
	for name, o := range mf.File {
		file := o[0]

		if err = buildDownload(file, basicPath, name, downloads); err != nil {
			return c.JSONResponse(schemas.NewErrors(err), 500)
		}
	}

	if err = builds.MarshalDownloads(downloads); err != nil {
		return c.JSONResponse(schemas.NewErrors(err), 500)
	}

	if err = controller.CreateDownload(project, version, unsafeConvert.StringToInt(build), builds); err != nil {
		return c.JSONResponse(schemas.NewErrors(err), 500)
	}

	return c.JSONResponse(schemas.NewResult("success"), 200)
}

func buildDownload(file *multipart.FileHeader, path, id string, downloads map[string]hssv2.ApplicationVersionsSchema) error {
	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()
	fileHash := fshash.NewReader(f, &fshash.Opt{Crypt: fshash.SHA256})
	if fileHash == "" {
		return errors.New("file hash is empty")
	}

	if len(downloads) != 0 {
		o, ok := downloads[id]
		if ok && o.Sha256 == fileHash {
			return nil
		}

		for _, version := range downloads {
			if version.Sha256 == fileHash {
				return nil
			}
		}
	}

	filePath := filepath.Join(path, file.Filename)
	if fsutil.IsExist(filePath) {
		_ = fsutil.Remove(filePath)
	}
	localFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer localFile.Close()
	_, err = io.Copy(localFile, f)
	if err != nil {
		return err
	}

	downloads[id] = hssv2.ApplicationVersionsSchema{
		Name:   file.Filename,
		Sha256: fileHash,
	}

	return nil
}
