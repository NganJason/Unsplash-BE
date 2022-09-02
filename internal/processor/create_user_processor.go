package processor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/NganJason/BE-template/internal"
	"github.com/NganJason/BE-template/internal/handler"
	"github.com/NganJason/BE-template/internal/model"
	"github.com/NganJason/BE-template/internal/vo"
	"github.com/NganJason/BE-template/pkg/cerr"
	"github.com/NganJason/BE-template/pkg/server"
)

func CreateUserProcessor(
	ctx context.Context,
	writer http.ResponseWriter,
	req *http.Request,
) *server.HandlerResp {
	request, ok := ctx.Value(internal.CtxRequestBody).(*vo.CreateUserRequest)
	if !ok {
		return server.NewHandlerResp(
			nil,
			cerr.New("assert request err", http.StatusBadRequest),
		)
	}

	response := &vo.CreateUserResponse{}

	p := &createUserProcessor{
		ctx:  ctx,
		req:  request,
		resp: response,
	}

	return p.process()
}

type createUserProcessor struct {
	ctx  context.Context
	req  *vo.CreateUserRequest
	resp *vo.CreateUserResponse
}

func (p *createUserProcessor) process() *server.HandlerResp {
	userDM := model.NewUserDM(p.ctx)

	h := handler.NewUserHandler(p.ctx, userDM)

	user, err := h.CreateUser(
		p.req,
	)
	if err != nil {
		return server.NewHandlerResp(
			nil,
			err,
		)
	}

	p.resp.User = user

	return server.NewHandlerResp(
		p.resp,
		nil,
	)
}

func (p *createUserProcessor) validateReq() error {
	if p.req.EmailAddress == nil || *p.req.EmailAddress == "" {
		return fmt.Errorf("email address cannot be empty")
	}

	if p.req.Password == nil || *p.req.Password == "" {
		return fmt.Errorf("password cannot be empty")
	}

	if p.req.Username == nil || *p.req.Username == "" {
		return fmt.Errorf("username cannot be empty")
	}

	return nil
}
