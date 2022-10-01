package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

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
		Addr:    ":" + GetPort(),
		Handler: handler,
	}

	clog.Info(context.Background(), fmt.Sprintf("Listening to port %s", GetPort()))

	server.ListenAndServe()
}

func GetPort() string {
	var port = os.Getenv("PORT")

	if port == "" {
		port = "8082"
	}

	return ":" + port
}
