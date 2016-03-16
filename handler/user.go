package handler

import (
	"github.com/banerwai/gommon/middleware"
	"github.com/martini-contrib/render"
)

func ShowSignup(ctx *middleware.Context, ren render.Render) {

	ren.HTML(200, "user/signup", ctx, render.HTMLOptions{
		Layout: "layout/layout",
	})
}

func ShowSignupEmployer(ctx *middleware.Context, ren render.Render) {

	ren.HTML(200, "user/signup_employer", ctx, render.HTMLOptions{
		Layout: "layout/layout",
	})
}

func ShowSignupContractor(ctx *middleware.Context, ren render.Render) {

	ren.HTML(200, "user/signup_contractor", ctx, render.HTMLOptions{
		Layout: "layout/layout",
	})
}

func ShowLogin(ctx *middleware.Context, ren render.Render) {

	ren.HTML(200, "user/login", ctx, render.HTMLOptions{
		Layout: "layout/layout",
	})
}
