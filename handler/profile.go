package handler

import (
	"fmt"
	"github.com/banerwai/gather/service"
	"github.com/banerwai/gommon/middleware"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"strconv"
	"time"
)

func SearchProfilesByKey(ctx *middleware.Context, ren render.Render, w http.ResponseWriter, r *http.Request) {

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

func ShowProfilesByCategory(ctx *middleware.Context, ren render.Render, params martini.Params) {
	_cat := params["category"]

	_category_id, _err := strconv.ParseInt(_cat, 10, 64)
	if _err != nil {
		ren.HTML(200, "profile/show_category", ctx, render.HTMLOptions{
			Layout: "layout/layout",
		})
		return
	}

	var _service service.ProfileService
	_profiles, _ := _service.GetProfilesByCategoryBean(_category_id, time.Now().Unix(), 10)

	ctx.Set("Profiles", _profiles)

	ren.HTML(200, "profile/show_category", ctx, render.HTMLOptions{
		Layout: "layout/layout",
	})
}

func ShowProfilesBySubCategory(ctx *middleware.Context, ren render.Render, params martini.Params) {
	_subcat := params["subcategory"]

	_subcategory_id, _err := strconv.ParseInt(_subcat, 10, 64)
	if _err != nil {
		ren.HTML(200, "profile/show_subcategory", ctx, render.HTMLOptions{
			Layout: "layout/layout",
		})
		return
	}

	var _service service.ProfileService
	_profiles, _ := _service.GetProfilesBySubCategoryBean(_subcategory_id, time.Now().Unix(), 10)

	ctx.Set("Profiles", _profiles)

	ren.HTML(200, "profile/show_subcategory", ctx, render.HTMLOptions{
		Layout: "layout/layout",
	})
}

func ShowProfileById(ctx *middleware.Context, ren render.Render, params martini.Params) {
	_profile_id := params["id"]
	var _service service.ProfileService
	_profile, _ := _service.GetProfileBean(_profile_id)

	ctx.Set("Profile", _profile)

	ren.HTML(200, "profile/show", ctx, render.HTMLOptions{
		Layout: "layout/layout",
	})
}
