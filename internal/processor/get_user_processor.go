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

func GetUserProcessor(
	ctx context.Context,
	writer http.ResponseWriter,
	req *http.Request,
) *server.HandlerResp {
	response := &vo.GetUserResponse{}

	request, ok := ctx.Value(internal.CtxRequestBody).(*vo.GetUserRequest)
	if !ok {
		return server.NewHandlerResp(
			response,
			cerr.New("assert request err", http.StatusBadRequest),
		)
	}

	p := &getUserProcessor{
		ctx:    ctx,
		writer: writer,
		req:    request,
		resp:   response,
	}

	return p.process()
}

type getUserProcessor struct {
	ctx    context.Context
	writer http.ResponseWriter
	req    *vo.GetUserRequest
	resp   *vo.GetUserResponse
}

func (p *getUserProcessor) process() *server.HandlerResp {
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
	h := handler.NewUserHandler(p.ctx, userDM)

	user, err := h.GetUser(
		p.req.UserID,
		nil,
		nil,
	)
	if err != nil {
		return server.NewHandlerResp(
			p.resp,
			err,
		)
	}

	p.resp.User = user

	return server.NewHandlerResp(
		p.resp,
		nil,
	)
}

func (p *getUserProcessor) validateReq() error {
	if p.req.UserID == nil || *p.req.UserID == 0 {
		return fmt.Errorf("userID cannot be empty")
	}

	return nil
}
