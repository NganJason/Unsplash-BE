package processor

import (
	"net/http"

	"github.com/NganJason/BE-template/internal/vo"
	"github.com/NganJason/BE-template/pkg/server"
)

func AllProcessors() []*server.Route {
	processors := []*server.Route{
		{
			Name:    vo.CmdHealthCheck,
			Method:  http.MethodGet,
			Path:    vo.PathHealthCheck,
			Handler: HealthCheck,
			Req:     nil,
		},
		{
			Name:    vo.CmdGetUser,
			Method:  http.MethodPost,
			Path:    vo.PathGetUser,
			Handler: GetUserProcessor,
			Req:     &vo.GetUserRequest{},
		},
		{
			Name:    vo.CmdCreateUser,
			Method:  http.MethodPost,
			Path:    vo.PathCreateUser,
			Handler: CreateUserProcessor,
			Req:     &vo.CreateUserRequest{},
		},
		{
			Name:    vo.CmdGetImages,
			Method:  http.MethodPost,
			Path:    vo.PathGetImages,
			Handler: GetImagesProcessor,
			Req:     &vo.GetImagesRequest{},
		},
		{
			Name:   vo.CmdUploadImage,
			Method: http.MethodPost,
			Path:   vo.PathUploadImage,
			Req:    &vo.UploadImageRequest{},
		},
		{
			Name:   vo.CmdLikeImage,
			Method: http.MethodPost,
			Path:   vo.PathLikeImage,
			Req:    &vo.LikeImageRequest{},
		},
		{
			Name:   vo.CmdDownloadImg,
			Method: http.MethodPost,
			Path:   vo.PathDownloadImg,
			Req:    &vo.DownloadImageRequest{},
		},
	}

	return processors
}
