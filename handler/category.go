package handler

import (
	"github.com/banerwai/gather/service"
	"github.com/banerwai/gommon/middleware"
	"github.com/martini-contrib/render"
)

func ShowCategories(ctx *middleware.Context, ren render.Render) {
	var _service service.CategoryService

	_categories := _service.GetCategoriesBean()

	ctx.Set("Categories", _categories)

	ren.HTML(200, "category/show_all", ctx, render.HTMLOptions{
		Layout: "layout/layout",
	})
}
