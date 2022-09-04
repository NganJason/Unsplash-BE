package processor

import (
	"context"
	"net/http"

	"github.com/NganJason/BE-template/internal"
	"github.com/NganJason/BE-template/internal/handler"
	"github.com/NganJason/BE-template/internal/model"
	"github.com/NganJason/BE-template/internal/service"
	"github.com/NganJason/BE-template/internal/util"
	"github.com/NganJason/BE-template/internal/vo"
	"github.com/NganJason/BE-template/pkg/cerr"
	"github.com/NganJason/BE-template/pkg/server"
)

func UploadImageProcessor(
	ctx context.Context,
	writer http.ResponseWriter,
	req *http.Request,
) *server.HandlerResp {
	request, ok := ctx.Value(internal.CtxRequestBody).(*vo.UploadImageRequest)
	if !ok {
		return server.NewHandlerResp(
			nil,
			cerr.New("assert request err", http.StatusBadRequest),
		)
	}

	response := &vo.UploadImageResponse{}

	userID, err := util.GetUserIDFromCookies(ctx)
	if err != nil {
		return server.NewHandlerResp(
			nil,
			cerr.New(
				err.Error(),
				http.StatusForbidden,
			),
		)
	}

	p := &uploadImageResponse{
		ctx:    ctx,
		req:    request,
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
	file := p.ctx.Value(internal.CtxFileBody)
	if file == nil {
		return server.NewHandlerResp(
			nil,
			cerr.New(
				"cannot parse file",
				http.StatusBadGateway,
			),
		)
	}

	fileBytes, _ := file.([]byte)
	if fileBytes == nil {
		return server.NewHandlerResp(
			nil,
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
		p.req.Desc,
	)
	if err != nil {
		return server.NewHandlerResp(
			nil,
			err,
		)
	}

	p.resp.Image = image

	return server.NewHandlerResp(
		p.resp,
		nil,
	)
}
