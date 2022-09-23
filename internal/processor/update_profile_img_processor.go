package processor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/NganJason/Unsplash-BE/internal"
	"github.com/NganJason/Unsplash-BE/internal/handler"
	"github.com/NganJason/Unsplash-BE/internal/model"
	"github.com/NganJason/Unsplash-BE/internal/service"
	"github.com/NganJason/Unsplash-BE/internal/util"
	"github.com/NganJason/Unsplash-BE/internal/vo"
	"github.com/NganJason/Unsplash-BE/pkg/cerr"
	"github.com/NganJason/Unsplash-BE/pkg/server"
)

func UpdateProfileImgProcessor(
	ctx context.Context,
	writer http.ResponseWriter,
	req *http.Request,
) *server.HandlerResp {
	response := &vo.UpdateProfileImgResponse{}

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

	file := ctx.Value(internal.CtxProfileImg)
	if file == nil {
		return server.NewHandlerResp(
			response,
			cerr.New(
				"cannot parse profile img file",
				http.StatusBadGateway,
			),
		)
	}

	fileBytes, _ := file.([]byte)
	if fileBytes == nil {
		return server.NewHandlerResp(
			response,
			cerr.New(
				"assert file error",
				http.StatusBadGateway,
			),
		)
	}

	p := &updateProfileImageProcessor{
		ctx:      ctx,
		resp:     response,
		userID:   userID,
		fileByte: fileBytes,
	}

	return p.process()
}

type updateProfileImageProcessor struct {
	ctx      context.Context
	req      *vo.UpdateProfileImgRequest
	resp     *vo.UpdateProfileImgResponse
	userID   *uint64
	fileByte []byte
}

func (p *updateProfileImageProcessor) process() *server.HandlerResp {
	userDM := model.NewUserDM(p.ctx)
	imageService, err := service.NewImageService(p.ctx)
	if err != nil {
		return server.NewHandlerResp(
			p.resp,
			cerr.New(
				fmt.Sprintf("init cloudinary err=%s", err.Error()),
				http.StatusBadGateway,
			),
		)
	}

	h := handler.NewUserHandler(
		p.ctx,
		userDM,
	)

	h.SetImageService(imageService)

	user, err := h.UpdateProfileImg(
		*p.userID,
		p.fileByte,
	)
	if err != nil {
		return server.NewHandlerResp(
			p.resp,
			err,
		)
	}

	p.resp.User = user

	return server.NewHandlerResp(
		p.resp,
		nil,
	)
}
