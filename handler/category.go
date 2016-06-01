package handler

import (
	"github.com/banerwai/gommon/middleware"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"github.com/banerwai/gather/service"
)

func ShowCategories(ctx *middleware.Context, ren render.Render) {
	var _service service.CategoryService

	_categories := _service.GetCategoriesBean()

	ctx.Set("Categories", _categories)

	ren.HTML(200, "category/show_all", ctx, render.HTMLOptions{
		Layout: "layout/layout",
	})
}

func ShowProfilesByCategory(ctx *middleware.Context, ren render.Render, params martini.Params) {
	_cat := params["category"]

	ctx.Set("Cat", _cat)

	ren.HTML(200, "profile/show_category", ctx, render.HTMLOptions{
		Layout: "layout/layout",
	})
}

func ShowProfilesBySubCategory(ctx *middleware.Context, ren render.Render, params martini.Params) {
	_cat := params["category"]
	_subcat := params["subcategory"]

	ctx.Set("Cat", _cat)
	ctx.Set("Subcat", _subcat)

	ren.HTML(200, "profile/show_subcategory", ctx, render.HTMLOptions{
		Layout: "layout/layout",
	})
}
