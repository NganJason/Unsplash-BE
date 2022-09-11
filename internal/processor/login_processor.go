package processor

import (
	"context"
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

func LoginProcessor(
	ctx context.Context,
	writer http.ResponseWriter,
	req *http.Request,
) *server.HandlerResp {
	response := &vo.LoginResponse{}

	request, ok := ctx.Value(internal.CtxRequestBody).(*vo.LoginRequest)
	if !ok {
		return server.NewHandlerResp(
			response,
			cerr.New("assert request err", http.StatusBadRequest),
		)
	}

	p := &loginProcessor{
		ctx:    ctx,
		writer: writer,
		req:    request,
		resp:   response,
	}

	return p.process()
}

type loginProcessor struct {
	ctx    context.Context
	writer http.ResponseWriter
	req    *vo.LoginRequest
	resp   *vo.LoginResponse
}

func (p *loginProcessor) process() *server.HandlerResp {
	userDM := model.NewUserDM(p.ctx)

	h := handler.NewUserHandler(p.ctx, userDM)

	user, err := h.GetUser(
		nil,
		p.req.EmailAddress,
		p.req.Password,
	)
	if err != nil {
		return server.NewHandlerResp(
			p.resp,
			err,
		)
	}

	cookie, err := util.GenerateCookies(
		strconv.FormatUint(*user.ID, 10),
	)
	if err != nil {
		return server.NewHandlerResp(
			p.resp,
			cerr.New(
				err.Error(),
				http.StatusBadGateway,
			),
		)
	}

	http.SetCookie(p.writer, cookie)

	p.resp.User = user

	return server.NewHandlerResp(
		p.resp,
		nil,
	)
}
