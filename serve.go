package main

import (
	"flag"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/bind"
)

func serve(config *Config) {
	if config.Server.Socket != "" {
		goji.ServeListener(bind.Socket(config.Server.Socket))
		return
	}

	if config.Server.Port != "" {
		flag.Set("bind", ":"+config.Server.Port)
	}
	goji.Serve()
}
