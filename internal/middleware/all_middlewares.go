package middleware

import "github.com/NganJason/BE-template/pkg/server"

func AllMiddlewares() []server.Middleware {
	middlewares := []server.Middleware{
		&ParseBodyMiddleware{},
	}

	return middlewares
}
