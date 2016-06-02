package main

import (
	"github.com/banerwai/gather/flagparse"
	"github.com/banerwai/gommon/middleware"
	"github.com/banerwai/website/usecases"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessionauth"
	"github.com/martini-contrib/sessions"
	"html/template"
	"log"
	"net/http"
)

func main() {

	defer log.Println("Shutdown complete!")

	_port := flagparse.BanerwaiWebPort
	log.Println("Starting Banerwai WebSite Listen on Port " + _port)

	store := sessions.NewCookieStore([]byte("BanerwaiSecret!!!"))
	m := martini.Classic()
	store.Options(sessions.Options{
		MaxAge: 0,
	})
	m.Use(sessions.Sessions("BanerwaiSession", store))
	m.Use(sessionauth.SessionUser(usecases.GenerateAnonymousUser))
	sessionauth.RedirectUrl = "/login"
	sessionauth.RedirectParam = "next"

	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Layout:     "layout/layout",
		Extensions: []string{".tmpl", ".html"},
		Charset:    "UTF-8",
		Funcs: []template.FuncMap{
			{
				"equal": func(args ...interface{}) bool {
					return args[0] == args[1]
				},
				"plus": func(args ...int) int {
					var result int
					for _, val := range args {
						result += val
					}
					return result
				},
			},
		},
	}))

	m.Use(HelperFuncs())

	// Setup static file serving
	m.Use(martini.Static("assets"))

	// Setup middleware.Context
	m.Use(middleware.InitContext())

	// Setup routing
	setupRoute(m)

	http.ListenAndServe(_port, m)
}
