package handler

import (
	"github.com/banerwai/gommon/middleware"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func ShowHowHelp(ctx *middleware.Context, ren render.Render, params martini.Params) {

	_info := params["help"]
	if len(_info) == 0 {
		ren.Redirect("/", 302)
		return
	}

	ctx.Set("Info", _info)

	ren.HTML(200, "how/help", ctx, render.HTMLOptions{
		Layout: "layout/layout",
	})
}
