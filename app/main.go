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

	server := &http.Server{
		Addr:    ":" + "8082",
		Handler: router,
	}

	clog.Info(context.Background(), "serving at 8082")

	server.ListenAndServe()
}
