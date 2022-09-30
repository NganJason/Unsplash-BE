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

func SearchUsersProcessor(
	ctx context.Context,
	writer http.ResponseWriter,
	req *http.Request,
) *server.HandlerResp {
	response := &vo.SearchUsersResponse{}

	request, ok := ctx.Value(internal.CtxRequestBody).(*vo.SearchUsersRequest)
	if !ok {
		return server.NewHandlerResp(
			response,
			cerr.New("assert request err", http.StatusBadRequest),
		)
	}

	p := &searchUsersProcessor{
		ctx:  ctx,
		req:  request,
		resp: response,
	}

	return p.process()
}

type searchUsersProcessor struct {
	ctx  context.Context
	req  *vo.SearchUsersRequest
	resp *vo.SearchUsersResponse
}

func (p *searchUsersProcessor) process() *server.HandlerResp {
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

	users, err := h.SearchUsers(
		*p.req.Keyword,
	)
	if err != nil {
		return server.NewHandlerResp(
			p.resp,
			err,
		)
	}

	p.resp.Users = users

	return server.NewHandlerResp(
		p.resp,
		nil,
	)
}

func (p *searchUsersProcessor) validateReq() error {
	if p.req.Keyword == nil || *p.req.Keyword == "" {
		return fmt.Errorf("keyword cannot be empty")
	}

	return nil
}
