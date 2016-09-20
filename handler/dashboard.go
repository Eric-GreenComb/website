package handler

import (
	"github.com/banerwai/gommon/middleware"
	"github.com/martini-contrib/render"
)

// DashboardIndexForm index form
func DashboardIndexForm(ctx *middleware.Context, ren render.Render) {

	ren.HTML(200, "dashboard/index", ctx, render.HTMLOptions{
		Layout: "layout/layout_dashboard",
	})
}
