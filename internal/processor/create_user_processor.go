package processor

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/NganJason/Unsplash-BE/internal"
	"github.com/NganJason/Unsplash-BE/internal/handler"
	"github.com/NganJason/Unsplash-BE/internal/model"
	"github.com/NganJason/Unsplash-BE/internal/util"
	"github.com/NganJason/Unsplash-BE/internal/vo"
	"github.com/NganJason/Unsplash-BE/pkg/cerr"
	"github.com/NganJason/Unsplash-BE/pkg/server"
)

func CreateUserProcessor(
	ctx context.Context,
	writer http.ResponseWriter,
	req *http.Request,
) *server.HandlerResp {
	response := &vo.CreateUserResponse{}

	request, ok := ctx.Value(internal.CtxRequestBody).(*vo.CreateUserRequest)
	if !ok {
		return server.NewHandlerResp(
			response,
			cerr.New("assert request err", http.StatusBadRequest),
		)
	}

	p := &createUserProcessor{
		ctx:    ctx,
		writer: writer,
		req:    request,
		resp:   response,
	}

	return p.process()
}

type createUserProcessor struct {
	ctx    context.Context
	writer http.ResponseWriter
	req    *vo.CreateUserRequest
	resp   *vo.CreateUserResponse
}

func (p *createUserProcessor) process() *server.HandlerResp {
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

	user, err := h.CreateUser(
		p.req,
	)
	if err != nil {
		return server.NewHandlerResp(
			p.resp,
			err,
		)
	}

	token, err := util.GenerateJWTToken(strconv.FormatUint(*user.ID, 10))
	if err != nil {
		return server.NewHandlerResp(
			p.resp,
			cerr.New(
				err.Error(),
				http.StatusBadGateway,
			),
		)
	}

	user.Token = util.StrPtr(token)
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
