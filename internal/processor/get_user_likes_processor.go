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

func GetUserLikesProcessor(
	ctx context.Context,
	writer http.ResponseWriter,
	req *http.Request,
) *server.HandlerResp {
	response := &vo.GetUserLikesResponse{}

	request, ok := ctx.Value(internal.CtxRequestBody).(*vo.GetUserLikesRequest)
	if !ok {
		return server.NewHandlerResp(
			response,
			cerr.New("assert request err", http.StatusBadRequest),
		)
	}

	p := &getUserLikesProcessor{
		ctx:    ctx,
		writer: writer,
		req:    request,
		resp:   response,
	}

	return p.process()
}

type getUserLikesProcessor struct {
	ctx    context.Context
	writer http.ResponseWriter
	req    *vo.GetUserLikesRequest
	resp   *vo.GetUserLikesResponse
}

func (p *getUserLikesProcessor) process() *server.HandlerResp {
	if err := p.validateReq(); err != nil {
		return server.NewHandlerResp(
			p.resp,
			cerr.New(
				err.Error(),
				http.StatusBadRequest,
			),
		)
	}

	userDM := model.NewUserDM(p.ctx)
	imageDM := model.NewImageDM(p.ctx)
	userLikeDM := model.NewUserLikeDM(p.ctx)

	h := handler.NewImageHandler(p.ctx, imageDM)
	h.SetUserDM(userDM)
	h.SetUserLikeDM(userLikeDM)
	

	images, nextCursor, err := h.GetImagesLikedByUser(
		p.req.UserID,
		p.req.PageSize,
		p.req.Cursor,
	)
	if err != nil {
		return server.NewHandlerResp(
			p.resp,
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

func (p *getUserLikesProcessor) validateReq() error {
	if p.req.UserID == nil || *p.req.UserID == 0 {
		return fmt.Errorf("userID cannot be empty")
	}

	if p.req.PageSize == nil || *p.req.PageSize == 0 {
		return fmt.Errorf("pageSize cannot be empty")
	}

	return nil
}
