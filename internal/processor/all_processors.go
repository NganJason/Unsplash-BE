package processor

import (
	"net/http"

	"github.com/NganJason/Unsplash-BE/internal/vo"
	"github.com/NganJason/Unsplash-BE/pkg/server"
)

func AllProcessors() []*server.Route {
	processors := []*server.Route{
		{
			Name:    vo.CmdHealthCheck,
			Method:  http.MethodPost,
			Path:    vo.PathHealthCheck,
			Handler: HealthCheck,
			Req:     vo.HealthCheckRequest{},
		},
		{
			Name:    vo.CmdLogin,
			Method:  http.MethodPost,
			Path:    vo.PathLogin,
			Handler: LoginProcessor,
			Req:     vo.LoginRequest{},
		},
		{
			Name:    vo.CmdLogout,
			Method:  http.MethodPost,
			Path:    vo.PathLogout,
			Handler: LogoutProcessor,
			Req:     nil,
		},
		{
			Name:    vo.CmdGetUser,
			Method:  http.MethodPost,
			Path:    vo.PathGetUser,
			Handler: GetUserProcessor,
			Req:     vo.GetUserRequest{},
		},
		{
			Name:    vo.CmdCreateUser,
			Method:  http.MethodPost,
			Path:    vo.PathCreateUser,
			Handler: CreateUserProcessor,
			Req:     vo.CreateUserRequest{},
		},
		{
			Name:    vo.CmdGetImages,
			Method:  http.MethodPost,
			Path:    vo.PathGetImages,
			Handler: GetImagesProcessor,
			Req:     vo.GetImagesRequest{},
		},
		{
			Name:    vo.CmdUploadImage,
			Method:  http.MethodPost,
			Path:    vo.PathUploadImage,
			Handler: UploadImageProcessor,
			Req:     nil,
		},
		{
			Name:    vo.CmdAddDeltaImage,
			Method:  http.MethodPost,
			Path:    vo.PathAddDeltaImage,
			Handler: AddDeltaImageProcessor,
			Req:     vo.AddDeltaImageRequest{},
		},
	}

	return processors
}
