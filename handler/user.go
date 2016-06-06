package handler

import (
	"errors"
	"fmt"
	"github.com/banerwai/gather/service"
	"github.com/banerwai/global/bean"
	"github.com/banerwai/gommon/middleware"
	"github.com/banerwai/gommon/regexp"
	"github.com/banerwai/website/usecases"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessionauth"
	"github.com/martini-contrib/sessions"
	"labix.org/v2/mgo/bson"
	"net/http"
	"strings"
)

func SignupUserForm(ctx *middleware.Context, ren render.Render) {

	ren.HTML(200, "user/signup", ctx, render.HTMLOptions{
		Layout: "layout/layout",
	})
}

func RegisterUser(ctx *middleware.Context, r *http.Request, ren render.Render) {
	r.ParseForm()
	_email, _pwd, _invited_email, _err := getSignupInfoByForm(r)

	if _err != nil {
		ren.Redirect("/signup/user", 302)
		return
	}
	var _user_service service.UserService

	var _user bean.User

	if len(_invited_email) == 0 {
		_user.Invited = bson.ObjectIdHex("5707cb10ae6faa1d1071a189")
	}
	_user.Email = _email
	_user.Pwd = _pwd

	fmt.Println(_user)
	_user_service.CreateBeanUser(_user)

	ren.Redirect("/", 302)
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

	ren.Redirect("/dashboard/index")
}

func getSignupInfoByForm(r *http.Request) (string, string, string, error) {
	_email := strings.TrimSpace(r.Form.Get("email"))
	_invited_email := strings.TrimSpace(r.Form.Get("invited_email"))
	_pwd := strings.TrimSpace(r.Form.Get("pwd"))
	_repwd := strings.TrimSpace(r.Form.Get("repwd"))

	if len(_email) == 0 || len(_pwd) == 0 || len(_repwd) == 0 {
		return "", "", "", errors.New("data can't be null")
	}

	if !regexp.IsEmail(_email) {
		return "", "", "", errors.New("input(email) can't match email regex")
	}
	if len(_invited_email) == 0 || !regexp.IsEmail(_invited_email) {
		_invited_email = ""
	}

	if _pwd != _repwd {
		return "", "", "", errors.New("pwd need equals repwd")
	}

	return _email, _pwd, _invited_email, nil
}
