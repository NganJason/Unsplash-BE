package processor

import (
	"context"
	"net/http"

	"github.com/NganJason/BE-template/internal"
	"github.com/NganJason/BE-template/internal/handler"
	"github.com/NganJason/BE-template/internal/model"
	"github.com/NganJason/BE-template/internal/util"
	"github.com/NganJason/BE-template/internal/vo"
	"github.com/NganJason/BE-template/pkg/cerr"
	"github.com/NganJason/BE-template/pkg/server"
)

func GetUserProcessor(
	ctx context.Context,
	writer http.ResponseWriter,
	req *http.Request,
) *server.HandlerResp {
	request, ok := ctx.Value(internal.CtxRequestBody).(*vo.GetUserRequest)
	if !ok {
		return server.NewHandlerResp(
			nil,
			cerr.New("assert request err", http.StatusBadRequest),
		)
	}

	response := &vo.GetUserResponse{}

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

	p := &getUserProcessor{
		ctx:    ctx,
		req:    request,
		resp:   response,
		userID: userID,
	}

	return p.process()
}

type getUserProcessor struct {
	ctx    context.Context
	req    *vo.GetUserRequest
	resp   *vo.GetUserResponse
	userID *uint64
}

func (p *getUserProcessor) process() *server.HandlerResp {
	userDM := model.NewUserDM(p.ctx)

	h := handler.NewUserHandler(p.ctx, userDM)

	user, err := h.GetUser(
		p.userID,
		p.req.EmailAddress,
		p.req.Password,
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
