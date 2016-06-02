package handler

import (
	"github.com/banerwai/gather/service"
	"github.com/banerwai/gommon/middleware"
	"github.com/banerwai/website/usecases"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessionauth"
	"github.com/martini-contrib/sessions"
	"net/http"
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

func LoginForm(ctx *middleware.Context, ren render.Render) {
	ren.HTML(200, "user/login", ctx, render.HTMLOptions{
		Layout: "layout/layout",
	})
}

func Logout(session sessions.Session, user sessionauth.User, ren render.Render) {
	sessionauth.Logout(session, user)
	ren.Redirect("/login")
}

func ValidateLogin(session sessions.Session, post_user usecases.UserModel, ren render.Render, req *http.Request) {
	// load user
	user := usecases.UserModel{}
	var _auth_service service.AuthService
	_user_dto, _err := _auth_service.LoginDto(post_user.Email, post_user.Password)
	if _err != nil {
		ren.Redirect(sessionauth.RedirectUrl)
		return
	}

	// authenticate the session
	user.Id = _user_dto.Id.Hex()
	user.Email = _user_dto.Email

	err := sessionauth.AuthenticateSession(session, &user)
	if err != nil {
		ren.JSON(500, err)
	}

	ren.Redirect("/")
}
