package processor

import (
	"context"
	"net/http"

	"github.com/NganJason/Unsplash-BE/internal"
	"github.com/NganJason/Unsplash-BE/internal/vo"
	"github.com/NganJason/Unsplash-BE/pkg/cerr"
	"github.com/NganJason/Unsplash-BE/pkg/server"
)

func TemplateProcessor(
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

	p := &templateProcessor{
		ctx:  ctx,
		req:  request,
		resp: response,
	}

	return p.process()
}

type templateProcessor struct {
	ctx  context.Context
	req  *vo.GetUserRequest
	resp *vo.GetUserResponse
}

func (p *templateProcessor) process() *server.HandlerResp {
	return nil
}
