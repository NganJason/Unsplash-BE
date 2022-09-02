package processor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/NganJason/BE-template/internal/vo"
	"github.com/NganJason/BE-template/pkg/server"
)

func HealthCheck(ctx context.Context, writer http.ResponseWriter, req *http.Request) *server.HandlerResp {
	// request := ctx.Value(internal.CtxRequestBody).(*vo.HealthCheckRequest)

	response := &vo.HealthCheckResponse{}
	response.Message = fmt.Sprintf("Echo from server")

	return server.NewHandlerResp(response, nil)
}
