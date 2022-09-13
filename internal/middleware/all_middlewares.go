package middleware

import (
	"github.com/NganJason/Unsplash-BE/internal/vo"
	"github.com/NganJason/Unsplash-BE/pkg/server"
)

func AllMiddlewares() []server.Middleware {
	parseBodyMiddleware := new(ParseBodyMiddleware)
	parseBodyMiddleware.Skip(vo.CmdUploadImage)

	parseFileBodyMiddleware := new(ParseFileBodyMiddleware)

	authMiddleware := new(AuthMiddleware)
	authMiddleware.Skip(
		vo.CmdLogin,
		vo.CmdLogout,
		vo.CmdCreateUser,
		vo.CmdHealthCheck,
		vo.CmdGetImages,
	)

	middlewares := []server.Middleware{
		parseBodyMiddleware,
		parseFileBodyMiddleware,
		authMiddleware,
	}

	return middlewares
}
