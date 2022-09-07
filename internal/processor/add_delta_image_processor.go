package processor

import (
	"context"
	"net/http"

	"github.com/NganJason/Unsplash-BE/internal"
	"github.com/NganJason/Unsplash-BE/internal/handler"
	"github.com/NganJason/Unsplash-BE/internal/model"
	"github.com/NganJason/Unsplash-BE/internal/util"
	"github.com/NganJason/Unsplash-BE/internal/vo"
	"github.com/NganJason/Unsplash-BE/pkg/cerr"
	"github.com/NganJason/Unsplash-BE/pkg/server"
)

func AddDeltaImageProcessor(
	ctx context.Context,
	writer http.ResponseWriter,
	req *http.Request,
) *server.HandlerResp {
	response := &vo.AddDeltaImageResponse{}

	request, ok := ctx.Value(internal.CtxRequestBody).(*vo.AddDeltaImageRequest)
	if !ok {
		return server.NewHandlerResp(
			response,
			cerr.New("assert request err", http.StatusBadRequest),
		)
	}

	userID, err := util.GetUserIDFromCookies(ctx)
	if err != nil {
		return server.NewHandlerResp(
			response,
			cerr.New(
				err.Error(),
				http.StatusForbidden,
			),
		)
	}

	p := &addDeltaImageProcessor{
		ctx:    ctx,
		req:    request,
		resp:   response,
		userID: userID,
	}

	return p.process()
}

type addDeltaImageProcessor struct {
	ctx    context.Context
	req    *vo.AddDeltaImageRequest
	resp   *vo.AddDeltaImageResponse
	userID *uint64
}

func (p *addDeltaImageProcessor) process() *server.HandlerResp {
	userDM := model.NewUserDM(p.ctx)
	imageDM := model.NewImageDM(p.ctx)
	userLikeDM := model.NewUserLikeDM(p.ctx)

	h := handler.NewUserHandler(p.ctx, userDM)
	h.SetImageDM(imageDM)
	h.SetUserLikeDM(userLikeDM)

	err := h.AddDeltaImage(*p.userID, *p.req.ImageID, &p.req.DeltaImage)
	if err != nil {
		return server.NewHandlerResp(
			p.resp,
			err,
		)
	}

	return server.NewHandlerResp(
		p.resp,
		nil,
	)
}
