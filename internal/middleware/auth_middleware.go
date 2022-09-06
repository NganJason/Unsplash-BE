package middleware

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
	"github.com/NganJason/Unsplash-BE/pkg/cookies"
	"github.com/NganJason/Unsplash-BE/pkg/server"
)

type AuthMiddleware struct {
	server.SkipMiddleware
	server.EmptyPostMiddleware
}

func (*AuthMiddleware) PreRequest(nextHandler server.Handler) server.Handler {
	return func(ctx context.Context, writer http.ResponseWriter, req *http.Request) *server.HandlerResp {
		c := cookies.ExtractCookie(req)
		if c == nil {
			return server.NewHandlerResp(
				&vo.CommonResponse{},
				cerr.New(
					"cookies not found, unauthorized",
					http.StatusUnauthorized,
				),
			)
		}

		jwt := c.Value
		auth, err := util.ParseJWTToken(jwt)
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

		user, err := h.GetUser(util.Uint64Ptr(userID), nil, nil)
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
