package middleware

import "github.com/NganJason/Unsplash-BE/pkg/server"

func AllMiddlewares() []server.Middleware {
	middlewares := []server.Middleware{
		&ParseBodyMiddleware{},
	}

	return middlewares
}
