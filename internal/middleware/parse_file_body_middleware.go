package middleware

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/NganJason/Unsplash-BE/internal"
	"github.com/NganJason/Unsplash-BE/internal/vo"
	"github.com/NganJason/Unsplash-BE/pkg/cerr"
	"github.com/NganJason/Unsplash-BE/pkg/server"
)

type ParseFileBodyMiddleware struct {
	server.SkipMiddleware
	server.EmptyPostMiddleware
}

func (*ParseFileBodyMiddleware) PreRequest(handler server.Handler) server.Handler {
	return func(ctx context.Context, writer http.ResponseWriter, req *http.Request) *server.HandlerResp {
		req.ParseMultipartForm(10 << 20)
		file, _, err := req.FormFile("img")
		if file == nil {
			return handler(ctx, writer, req)
		}

		defer file.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			return server.NewHandlerResp(
				&vo.CommonResponse{},
				cerr.New(
					fmt.Sprintf("read fileBytes err=%s", err.Error()),
					http.StatusBadGateway,
				),
			)
		}

		ctx = context.WithValue(
			req.Context(),
			internal.CtxFormBodyImg,
			fileBytes,
		)

		return handler(ctx, writer, req)
	}
}
