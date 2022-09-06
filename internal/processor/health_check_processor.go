package processor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/NganJason/Unsplash-BE/internal"
	"github.com/NganJason/Unsplash-BE/internal/vo"
	"github.com/NganJason/Unsplash-BE/pkg/server"
)

func HealthCheck(ctx context.Context, writer http.ResponseWriter, req *http.Request) *server.HandlerResp {
	request := ctx.Value(internal.CtxRequestBody).(*vo.HealthCheckRequest)

	response := &vo.HealthCheckResponse{}
	response.Message = fmt.Sprintf("Echo from server=%s", request.Message)

	return server.NewHandlerResp(response, nil)
}
