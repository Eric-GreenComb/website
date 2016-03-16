package main

import (
	"github.com/banerwai/gommon/middleware"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"html/template"
)

func main() {

	store := sessions.NewCookieStore([]byte("BanerwaiSecret!!!"))
	m := martini.Classic()
	store.Options(sessions.Options{
		MaxAge: 0,
	})
	m.Use(sessions.Sessions("BanerwaiSession", store))

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

	// Setup static file serving
	m.Use(martini.Static("assets"))

	// Setup middleware.Context
	m.Use(middleware.InitContext())

	// Setup routing
	setupRoute(m)

	m.Run()
}
