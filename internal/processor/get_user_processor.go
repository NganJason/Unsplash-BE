package processor

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/NganJason/Unsplash-BE/internal/handler"
	"github.com/NganJason/Unsplash-BE/internal/model"
	"github.com/NganJason/Unsplash-BE/internal/util"
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

	userID, err := util.GetUserIDFromCookies(ctx)
	if err != nil {
		return server.NewHandlerResp(
			response,
			cerr.New(
				fmt.Sprintf("parse cookie err=%s", err.Error()),
				http.StatusUnauthorized,
			),
		)
	}

	if userID == nil {
		return server.NewHandlerResp(
			response,
			cerr.New(
				"userID is nil in cookies",
				http.StatusUnauthorized,
			),
		)
	}

	p := &getUserProcessor{
		ctx:    ctx,
		writer: writer,
		resp:   response,
		userID: userID,
	}

	return p.process()
}

type getUserProcessor struct {
	ctx    context.Context
	writer http.ResponseWriter
	req    *vo.GetUserRequest
	resp   *vo.GetUserResponse
	userID *uint64
}

func (p *getUserProcessor) process() *server.HandlerResp {
	userDM := model.NewUserDM(p.ctx)

	h := handler.NewUserHandler(p.ctx, userDM)

	user, err := h.GetUser(
		p.userID,
		nil,
		nil,
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
