package handler

import (
	"github.com/banerwai/gommon/middleware"
	"github.com/martini-contrib/render"
)

// ShowPage404 show 404 page
func ShowPage404(ctx *middleware.Context, ren render.Render) {

	ren.HTML(200, "system/404", ctx, render.HTMLOptions{
		Layout: "layout/layout",
	})
}
