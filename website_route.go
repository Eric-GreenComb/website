package main

import (
	"github.com/banerwai/website/handler"
	"github.com/go-martini/martini"
)

func setupRoute(m *martini.ClassicMartini) {
	setupSystemRoute(m)

	setupCategoryRoute(m)

	setupProfileRoute(m)

	setupUserRoute(m)

	setupInfoRoute(m)

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
	m.Group("/i/catogery", func(r martini.Router) {
		r.Get("", handler.ShowCategories)
	})

	m.Group("/cat", func(r martini.Router) {
		r.Get("/:short", handler.ShowCategoryIndex)
	})
}

func setupProfileRoute(m *martini.ClassicMartini) {
	m.Group("/o/profiles/browse", func(r martini.Router) {
		r.Get("", handler.ShowProfilesByKey)
		r.Get("/c/:category", handler.ShowProfilesByCategory)
	})
}

func setupUserRoute(m *martini.ClassicMartini) {
	m.Group("/signup", func(r martini.Router) {
		r.Get("", handler.ShowSignup)
		r.Get("/employer", handler.ShowSignupEmployer)
		r.Get("/contractor", handler.ShowSignupContractor)
	})

	m.Group("/login", func(r martini.Router) {
		r.Get("", handler.ShowLogin)
	})
}

func setup404Route(m *martini.ClassicMartini) {
	m.NotFound(handler.ShowPage404)
}
