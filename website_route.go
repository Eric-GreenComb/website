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
		r.Get("/:id", handler.ShowProfileById)
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
