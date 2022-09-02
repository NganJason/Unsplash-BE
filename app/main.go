package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/NganJason/BE-template/internal/config"
	"github.com/NganJason/BE-template/internal/middleware"
	"github.com/NganJason/BE-template/internal/processor"
	"github.com/NganJason/BE-template/pkg/clog"
	"github.com/NganJason/BE-template/pkg/server"
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
