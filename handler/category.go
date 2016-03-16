package handler

import (
	"github.com/banerwai/gommon/middleware"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func ShowCategories(ctx *middleware.Context, ren render.Render) {

	ren.HTML(200, "category/show_all", ctx, render.HTMLOptions{
		Layout: "layout/layout",
	})
}

func ShowCategoryIndex(ctx *middleware.Context, ren render.Render, params martini.Params) {

	_cat := params["short"]
	if len(_cat) == 0 {
		ren.Redirect("/", 302)
		return
	}
	ctx.Set("Cat", _cat)
	ren.HTML(200, "category/show_cat", ctx, render.HTMLOptions{
		Layout: "layout/layout",
	})
}
