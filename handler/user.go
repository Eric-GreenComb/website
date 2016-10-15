package handler

import (
	"errors"
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

// SignupUserForm signup page
func SignupUserForm(ctx *middleware.Context, ren render.Render) {

	ren.HTML(200, "user/signup", ctx, render.HTMLOptions{
		Layout: "layout/layout",
	})
}

// RegisterUser register page
func RegisterUser(ctx *middleware.Context, r *http.Request, ren render.Render) {
	r.ParseForm()
	_email, _pwd, _invitedEmail, _err := getSignupInfoByForm(r)

	if _err != nil {
		ren.Redirect("/signup/user", 302)
		return
	}
	var _userService service.UserService

	var _user bean.User

	if len(_invitedEmail) == 0 {
		_user.Invited = bson.ObjectIdHex("5707cb10ae6faa1d1071a189")
	}
	_user.Email = _email
	_user.Pwd = _pwd

	_userService.CreateBeanUser(_user)

	ren.Redirect("/", 302)
}

// LoginForm login page
func LoginForm(ctx *middleware.Context, ren render.Render) {
	ren.HTML(200, "user/login", ctx, render.HTMLOptions{
		Layout: "layout/layout",
	})
}

// Logout logout handler
func Logout(session sessions.Session, user sessionauth.User, ren render.Render) {
	sessionauth.Logout(session, user)
	ren.Redirect("/login")
}

// ValidateLogin validate user
func ValidateLogin(session sessions.Session, postUser usecases.UserModel, ren render.Render, req *http.Request) {
	// load user
	user := usecases.UserModel{}
	var _authService service.AuthService
	_userDto, _err := _authService.LoginDto(postUser.Email, postUser.Password)
	if _err != nil {
		ren.Redirect(sessionauth.RedirectUrl)
		return
	}

	// authenticate the session
	user.Id = _userDto.ID.Hex()
	user.Email = _userDto.Email

	err := sessionauth.AuthenticateSession(session, &user)
	if err != nil {
		ren.JSON(500, err)
	}

	ren.Redirect("/dashboard/index")
}

func getSignupInfoByForm(r *http.Request) (string, string, string, error) {
	_email := strings.TrimSpace(r.Form.Get("email"))
	_invitedEmail := strings.TrimSpace(r.Form.Get("invited_email"))
	_pwd := strings.TrimSpace(r.Form.Get("pwd"))
	_repwd := strings.TrimSpace(r.Form.Get("repwd"))

	if len(_email) == 0 || len(_pwd) == 0 || len(_repwd) == 0 {
		return "", "", "", errors.New("data can't be null")
	}

	if !regexp.IsEmail(_email) {
		return "", "", "", errors.New("input(email) can't match email regex")
	}
	if len(_invitedEmail) == 0 || !regexp.IsEmail(_invitedEmail) {
		_invitedEmail = ""
	}

	if _pwd != _repwd {
		return "", "", "", errors.New("pwd need equals repwd")
	}

	return _email, _pwd, _invitedEmail, nil
}
