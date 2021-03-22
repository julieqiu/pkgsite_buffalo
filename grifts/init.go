package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/julieqiu/pkgsite_buffalo/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
