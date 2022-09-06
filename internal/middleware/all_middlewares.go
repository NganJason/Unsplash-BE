package middleware

import (
	"github.com/NganJason/Unsplash-BE/internal/vo"
	"github.com/NganJason/Unsplash-BE/pkg/server"
)

func AllMiddlewares() []server.Middleware {
	parseBodyMiddleware := &ParseBodyMiddleware{}
	parseBodyMiddleware.Skip(vo.CmdUploadImage)

	middlewares := []server.Middleware{
		parseBodyMiddleware,
		&ParseFileBodyMiddleware{},
	}

	return middlewares
}
