package processor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/NganJason/Unsplash-BE/internal"
	"github.com/NganJason/Unsplash-BE/internal/handler"
	"github.com/NganJason/Unsplash-BE/internal/model"
	"github.com/NganJason/Unsplash-BE/internal/vo"
	"github.com/NganJason/Unsplash-BE/pkg/cerr"
	"github.com/NganJason/Unsplash-BE/pkg/server"
)

func GetImagesProcessor(
	ctx context.Context,
	writer http.ResponseWriter,
	req *http.Request,
) *server.HandlerResp {
	request, ok := ctx.Value(internal.CtxRequestBody).(*vo.GetImagesRequest)
	if !ok {
		return server.NewHandlerResp(
			nil,
			cerr.New("assert request err", http.StatusBadRequest),
		)
	}

	response := &vo.GetImagesResponse{}

	p := &getImagesProcessor{
		ctx:  ctx,
		req:  request,
		resp: response,
	}

	return p.process()
}

type getImagesProcessor struct {
	ctx  context.Context
	req  *vo.GetImagesRequest
	resp *vo.GetImagesResponse
}

func (p *getImagesProcessor) process() *server.HandlerResp {
	if err := p.validateReq(); err != nil {
		return server.NewHandlerResp(
			nil,
			cerr.New(
				err.Error(),
				http.StatusBadRequest,
			),
		)
	}

	imageDM := model.NewImageDM(p.ctx)
	userDM := model.NewUserDM(p.ctx)

	h := handler.NewImageHandler(
		p.ctx,
		imageDM,
	)
	h.SetUserDM(userDM)

	images, nextCursor, err := h.GetImages(
		p.req.PageSize,
		p.req.Cursor,
	)
	if err != nil {
		return server.NewHandlerResp(
			nil,
			err,
		)
	}

	p.resp.Images = images
	p.resp.NextCursor = nextCursor

	return server.NewHandlerResp(
		p.resp,
		nil,
	)
}

func (p *getImagesProcessor) validateReq() error {
	if p.req.PageSize == nil || *p.req.PageSize == 0 {
		return fmt.Errorf("page size cannot be empty")
	}
	return nil
}
