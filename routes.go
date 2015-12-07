package main

import (
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

func routes() {
	apiRoutes()
}

func apiRoutes() *web.Mux {
	api := web.New()
	goji.Handle("/api/*", api)

	api.Use(middleware.SubRouter)

	api.Get("/tables", tableList)
	api.Get("/tables/:tablename", tableDetail)

	return api
}
