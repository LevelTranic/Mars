package _default

import (
	"github.com/savsgio/atreugo/v11"

	hssv2 "Mars/shared/httpschemas/v2"
)

func Index(c *atreugo.RequestCtx) error { return c.JSONResponse(hssv2.DefaultIndex, 200) }
