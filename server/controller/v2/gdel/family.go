package gdel

import (
	"github.com/savsgio/atreugo/v11"

	"Mars/server/helper"
)

func Family(c *atreugo.RequestCtx) error {
	defer func() {
		if p := recover(); p != nil {
			_ = helper.HandleInternalError(c, p)
			return
		}
	}()
	return nil
}
