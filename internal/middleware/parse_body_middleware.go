package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/NganJason/Unsplash-BE/internal"
	"github.com/NganJason/Unsplash-BE/internal/vo"
	"github.com/NganJason/Unsplash-BE/pkg/cerr"
	"github.com/NganJason/Unsplash-BE/pkg/server"
)

type ParseBodyMiddleware struct {
	server.SkipMiddleware
	server.EmptyPostMiddleware
}

func (*ParseBodyMiddleware) PreRequest(handler server.Handler) server.Handler {
	return func(ctx context.Context, writer http.ResponseWriter, req *http.Request) *server.HandlerResp {
		requestStruct := ctx.Value(server.CtxRequestStruct)
		var reqPayload interface{}
		body, _ := ioutil.ReadAll(req.Body)

		if len(body) > 0 || requestStruct != nil {
			if requestStruct != nil {
				reqPayload = reflect.New(reflect.TypeOf(requestStruct)).Interface()
			} else {
				reqPayload = map[string]interface{}{}
			}

			d := json.NewDecoder(bytes.NewReader(body))
			d.UseNumber()

			if err := d.Decode(&reqPayload); err != nil {
				return server.NewHandlerResp(
					&vo.CommonResponse{},
					cerr.New(
						fmt.Sprintf("decode req payload err=%s", err.Error()),
						http.StatusBadRequest,
					),
				)
			}
			ctx = context.WithValue(ctx, internal.CtxRequestBody, reqPayload)
		}

		return handler(ctx, writer, req)
	}
}
