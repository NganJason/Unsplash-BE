package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/NganJason/Unsplash-BE/internal/handler"
	"github.com/NganJason/Unsplash-BE/internal/model"
	"github.com/NganJason/Unsplash-BE/internal/util"
	"github.com/NganJason/Unsplash-BE/internal/vo"
	"github.com/NganJason/Unsplash-BE/pkg/cerr"
	"github.com/NganJason/Unsplash-BE/pkg/cookies"
	"github.com/NganJason/Unsplash-BE/pkg/server"
)

type AuthMiddleware struct {
	server.SkipMiddleware
	server.EmptyPostMiddleware
}

func (*AuthMiddleware) PreRequest(nextHandler server.Handler) server.Handler {
	return func(ctx context.Context, writer http.ResponseWriter, req *http.Request) *server.HandlerResp {
		reqToken := req.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")

		var token string
		if len(splitToken) > 1 {
			token = splitToken[1]
		} else {
			return server.NewHandlerResp(
				&vo.CommonResponse{},
				cerr.New(
					"token not found, unauthorized",
					http.StatusUnauthorized,
				),
			)
		}

		auth, err := util.ParseJWTToken(token)
		if err != nil || auth == nil {
			return server.NewHandlerResp(
				&vo.CommonResponse{},
				cerr.New(
					fmt.Sprintf("parse jwt token err=%s", err.Error()),
					http.StatusBadGateway,
				),
			)
		}

		if auth.Valid() != nil {
			return server.NewHandlerResp(
				&vo.CommonResponse{},
				cerr.New(
					fmt.Sprintf("invalid token err=%s", auth.Valid().Error()),
					http.StatusUnauthorized,
				),
			)
		}

		userIDStr := auth.Value
		userID, err := strconv.ParseUint(userIDStr, 0, 64)
		if err != nil {
			return server.NewHandlerResp(
				&vo.CommonResponse{},
				cerr.New(
					fmt.Sprintf("parse userIDStr err=%s", err.Error()),
					http.StatusBadGateway,
				),
			)
		}

		userDM := model.NewUserDM(req.Context())
		h := handler.NewUserHandler(req.Context(), userDM)

		user, err := h.GetUser(util.Uint64Ptr(userID), nil, nil, nil)
		if err != nil {
			return server.NewHandlerResp(
				&vo.CommonResponse{},
				err,
			)
		}

		if user == nil {
			return server.NewHandlerResp(
				&vo.CommonResponse{},
				cerr.New(
					"invalid userID",
					http.StatusUnauthorized,
				),
			)
		}

		ctx = cookies.AddClientCookieValToCtx(
			ctx,
			&userIDStr,
		)

		return nextHandler(ctx, writer, req)
	}
}
