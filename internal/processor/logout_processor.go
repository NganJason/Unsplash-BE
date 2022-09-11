package processor

import (
	"context"
	"net/http"

	"github.com/NganJason/Unsplash-BE/internal/vo"
	"github.com/NganJason/Unsplash-BE/pkg/cookies"
	"github.com/NganJason/Unsplash-BE/pkg/server"
)

func LogoutProcessor(
	ctx context.Context,
	writer http.ResponseWriter,
	req *http.Request,
) *server.HandlerResp {
	deleteCookie := cookies.DeleteCookie()
	http.SetCookie(writer, deleteCookie)

	return server.NewHandlerResp(
		&vo.CommonResponse{},
		nil,
	)
}
