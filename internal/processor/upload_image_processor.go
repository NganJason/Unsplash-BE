package processor

import (
	"context"
	"net/http"

	"github.com/NganJason/Unsplash-BE/internal"
	"github.com/NganJason/Unsplash-BE/internal/handler"
	"github.com/NganJason/Unsplash-BE/internal/model"
	"github.com/NganJason/Unsplash-BE/internal/service"
	"github.com/NganJason/Unsplash-BE/internal/util"
	"github.com/NganJason/Unsplash-BE/internal/vo"
	"github.com/NganJason/Unsplash-BE/pkg/cerr"
	"github.com/NganJason/Unsplash-BE/pkg/server"
)

func UploadImageProcessor(
	ctx context.Context,
	writer http.ResponseWriter,
	req *http.Request,
) *server.HandlerResp {
	response := &vo.UploadImageResponse{}

	userID, err := util.GetUserIDFromCookies(ctx)
	if err != nil {
		return server.NewHandlerResp(
			response,
			cerr.New(
				err.Error(),
				http.StatusForbidden,
			),
		)
	}

	p := &uploadImageResponse{
		ctx:    ctx,
		resp:   response,
		userID: userID,
	}

	return p.process()
}

type uploadImageResponse struct {
	ctx    context.Context
	req    *vo.UploadImageRequest
	resp   *vo.UploadImageResponse
	userID *uint64
}

func (p *uploadImageResponse) process() *server.HandlerResp {
	file := p.ctx.Value(internal.CtxFormBodyImg)
	if file == nil {
		return server.NewHandlerResp(
			p.resp,
			cerr.New(
				"cannot parse file",
				http.StatusBadGateway,
			),
		)
	}

	fileBytes, _ := file.([]byte)
	if fileBytes == nil {
		return server.NewHandlerResp(
			p.resp,
			cerr.New(
				"assert file error",
				http.StatusBadGateway,
			),
		)
	}

	imageDM := model.NewImageDM(p.ctx)
	userDM := model.NewUserDM(p.ctx)
	imageService := service.NewImageService(p.ctx)

	h := handler.NewImageHandler(
		p.ctx,
		imageDM,
	)
	h.SetUserDM(userDM)
	h.SetImageService(imageService)

	image, err := h.UploadImage(
		fileBytes,
		*p.userID,
		util.StrPtr(""),
	)
	if err != nil {
		return server.NewHandlerResp(
			p.resp,
			err,
		)
	}

	p.resp.Image = image

	return server.NewHandlerResp(
		p.resp,
		nil,
	)
}
