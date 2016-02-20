package main

import (
	"github.com/banerwai/website/handler"
	"github.com/go-martini/martini"
)

func setupSystemRoute(m *martini.ClassicMartini) {
	m.Group("/", func(r martini.Router) {
		r.Get("", handler.ShowIndex)

		r.Get("ping", func() string {
			return "pong"
		})
	})
}
