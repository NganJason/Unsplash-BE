package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/NganJason/Unsplash-BE/internal/config"
	"github.com/NganJason/Unsplash-BE/internal/middleware"
	"github.com/NganJason/Unsplash-BE/internal/processor"
	"github.com/NganJason/Unsplash-BE/pkg/clog"
	"github.com/NganJason/Unsplash-BE/pkg/server"
	"github.com/rs/cors"
)

func main() {
	clog.SetMinLogLevel(clog.LevelInfo)

	err := config.InitGlobalConfig()
	if err != nil {
		clog.Fatal(
			context.Background(),
			fmt.Sprintf("init config err=%s", err.Error()),
		)
	}

	router := server.NewMainRouter()

	for _, middleware := range middleware.AllMiddlewares() {
		router.AddMiddleware(middleware)
	}

	for _, route := range processor.AllProcessors() {
		router.AddRoute(route)
	}

	c := cors.New(
		cors.Options{
			AllowedOrigins:   []string{"http://localhost:3000"},
			AllowCredentials: true,
			AllowedMethods:   []string{"POST", "GET", "OPTIONS"},
		},
	)

	handler := c.Handler(router)

	server := &http.Server{
		Addr:    ":" + "8082",
		Handler: handler,
	}

	clog.Info(context.Background(), "serving at 8082")

	server.ListenAndServe()
}
