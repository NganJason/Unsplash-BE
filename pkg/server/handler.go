package server

import (
	"context"
	"net/http"
)

type Handler func(ctx context.Context, writer http.ResponseWriter, req *http.Request) *HandlerResp
type RespHandler func(ctx context.Context, writer http.ResponseWriter, req *http.Request, resp *HandlerResp) *HandlerResp
