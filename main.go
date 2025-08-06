package main

import (
	"go-starter/app/providers"
	"go-starter/core/server"
	"net/http"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		providers.Init(),
		fx.Provide(server.NewHTTPServer),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
