package handler

import (
	"fmt"
	"github.com/banerwai/gommon/middleware"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

func ShowProfilesByKey(ctx *middleware.Context, ren render.Render, w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	_q := r.Form.Get("q")

	if len(_q) == 0 {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, "Need url key : /o/profiles/browse/?q=key")
		return
	}

	ctx.Set("key", _q)

	ren.HTML(200, "profile/search_key", ctx, render.HTMLOptions{
		Layout: "layout/layout",
	})
}

func SearchProfilesBySubCategory(ctx *middleware.Context, ren render.Render, w http.ResponseWriter, r *http.Request, params martini.Params) {
	fmt.Println("SearchProfilesBySubCategory")
	_subcat := params["subcategory"]

	r.ParseForm()
	_pt := r.Form.Get("pt")

	ctx.Set("Subcat", _subcat)
	ctx.Set("pt", _pt)

	ren.HTML(200, "profile/search_subcategory", ctx, render.HTMLOptions{
		Layout: "layout/layout",
	})
}

func ShowProfileById(ctx *middleware.Context, ren render.Render) {
	ren.HTML(200, "profile/show", ctx, render.HTMLOptions{
		Layout: "layout/layout",
	})
}
