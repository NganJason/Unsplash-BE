package middleware

import (
	"github.com/NganJason/Unsplash-BE/internal/vo"
	"github.com/NganJason/Unsplash-BE/pkg/server"
)

func AllMiddlewares() []server.Middleware {
	parseBodyMiddleware := new(ParseBodyMiddleware)
	parseBodyMiddleware.Skip(
		vo.CmdUploadImage,
		vo.CmdSeedData,
		vo.CmdUpdateProfileImg,
	)

	parseFileBodyMiddleware := new(ParseFileBodyMiddleware)
	

	authMiddleware := new(AuthMiddleware)
	authMiddleware.Skip(
		vo.CmdLogin,
		vo.CmdLogout,
		vo.CmdCreateUser,
		vo.CmdHealthCheck,
		vo.CmdGetImages,
		vo.CmdSeedData,
		vo.CmdGetUser,
		vo.CmdGetUserLikes,
	)

	middlewares := []server.Middleware{
		parseBodyMiddleware,
		parseFileBodyMiddleware,
		authMiddleware,
	}

	return middlewares
}
