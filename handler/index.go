package handler

import (
	"github.com/banerwai/gommon/middleware"
	"github.com/martini-contrib/render"
)

// ShowIndex show index page
func ShowIndex(ctx *middleware.Context, ren render.Render) {

	ctx.Set("Website", "班儿外")
	ctx.Set("WebsiteTitle", "A Banerwai Website")
	ctx.Set("WebsiteDetail", "This is a website for banerwai.")

	ren.HTML(200, "index/index", ctx, render.HTMLOptions{
		Layout: "layout/layout_index",
	})
}
