package main

import (
	"github.com/banerwai/website/handler"
	"github.com/banerwai/website/usecases"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/sessionauth"
)

func setupRoute(m *martini.ClassicMartini) {
	setupSystemRoute(m)

	setupCategoryRoute(m)

	setupProfileRoute(m)

	setupUserRoute(m)

	setupInfoRoute(m)

	setupDashboardRoute(m)

	// 404
	setup404Route(m)
}

func setupSystemRoute(m *martini.ClassicMartini) {
	m.Group("/", func(r martini.Router) {
		r.Get("", handler.ShowIndex)

		r.Get("ping", func() string {
			return "pong"
		})
	})
}

func setupInfoRoute(m *martini.ClassicMartini) {
	m.Group("/i/how", func(r martini.Router) {
		r.Get("/:help", handler.ShowHowHelp)
	})
}

func setupCategoryRoute(m *martini.ClassicMartini) {
	m.Group("/c", func(r martini.Router) {
		r.Get("/", handler.ShowCategories)
	})
}

func setupProfileRoute(m *martini.ClassicMartini) {
	m.Group("/o/profiles/browse", func(r martini.Router) {
		r.Get("/", handler.SearchProfilesByKey)

	})

	m.Group("/o/profiles/c", func(r martini.Router) {
		r.Get("/:category", handler.ShowProfilesByCategory)
		r.Get("/:category/sc/:subcategory", handler.ShowProfilesBySubCategory)
	})

	m.Group("/freelancer", func(r martini.Router) {
		r.Get("/:id", handler.ShowProfileByID)
	})
}

func setupUserRoute(m *martini.ClassicMartini) {
	m.Group("/signup", func(r martini.Router) {
		r.Get("/user", handler.SignupUserForm)
		r.Post("/user", handler.RegisterUser)
		// r.Get("/Invited/:invited", handler.SignupInvitedForm)
		// r.Post("/Invited/:invited", handler.RegisterInvited)
	})

	m.Group("/", func(r martini.Router) {
		r.Get("login", handler.LoginForm)
		r.Post("login", binding.Bind(usecases.UserModel{}), handler.ValidateLogin)
		r.Get("logout", sessionauth.LoginRequired, handler.Logout)
	})
}

func setupDashboardRoute(m *martini.ClassicMartini) {
	m.Group("/dashboard", func(r martini.Router) {
		r.Get("/index", sessionauth.LoginRequired, handler.DashboardIndexForm)
	})
}

func setup404Route(m *martini.ClassicMartini) {
	m.NotFound(handler.ShowPage404)
}
