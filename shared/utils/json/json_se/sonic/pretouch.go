//go:build amd64

package sonic

import (
	"fmt"
	"log/slog"
	"reflect"
	"time"

	"github.com/3JoB/ulib/timings"
	"github.com/bytedance/sonic"

	dbsv2 "Mars/database/v2/schemas"
	schemas2 "Mars/server/schemas"
	"Mars/shared/configure"
	hssv2 "Mars/shared/httpschemas/v2"
)

func PreTouch() {
	if configure.Get().Json.Runtime != "sonic" || !configure.Get().Json.Sonic.JITPretouch {
		return
	}
	timer := time.Now()
	errors := 0
	var b = []reflect.Type{
		// b
		reflect.TypeOf(dbsv2.Change{}),
		reflect.TypeOf([]dbsv2.Change{}),
		reflect.TypeOf(hssv2.NewDownloadSchema{}),
		reflect.TypeOf(hssv2.IndexVersion{}),
		reflect.TypeOf(hssv2.ApplicationVersionsSchema{}),
		reflect.TypeOf(map[string]hssv2.ApplicationVersionsSchema{}),

		// c
		reflect.TypeOf(schemas2.ProjectRootSchema{}),
		reflect.TypeOf(schemas2.ResultSchema{}),
		reflect.TypeOf(schemas2.ErrorSchema{}),
		reflect.TypeOf(schemas2.NewBuildSchema{}),
		reflect.TypeOf(schemas2.NewVersionGroupSchema{}),
		reflect.TypeOf(schemas2.NewVersionSchema{}),
		reflect.TypeOf(schemas2.NewProjectSchema{}),
		reflect.TypeOf(schemas2.NewBuildResult{}),
		reflect.TypeOf(schemas2.BuildsSchema{}),
		reflect.TypeOf([]schemas2.BuildsSchema{}),
		reflect.TypeOf(schemas2.FamilyVersionsSchema{}),
		reflect.TypeOf(schemas2.FamilyVersionsBuildsSchema{}),
		reflect.TypeOf(schemas2.ProjectVersionBuildSchema{}),
		reflect.TypeOf(schemas2.ProjectVersionsBuildsSchema{}),
		reflect.TypeOf(schemas2.ProjectVersionsSchema{}),
		reflect.TypeOf(schemas2.ProjectSchema{}),
		reflect.TypeOf(schemas2.ProjectsSchema{}),
	}
	timings.ParallelForEach(b, timings.AttrSlice(b), func(v int, b reflect.Type) {
		if err := sonic.Pretouch(b); err != nil {
			errors++
		}
	})
	slog.Info(fmt.Sprintf("JITPretouch completed, compilation failed %v times, took %v ms.", errors, time.Since(timer).Milliseconds()))
}
