package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NganJason/Unsplash-BE/internal"
	"github.com/NganJason/Unsplash-BE/internal/handler"
	"github.com/NganJason/Unsplash-BE/internal/model"
	"github.com/NganJason/Unsplash-BE/internal/service"
	"github.com/NganJason/Unsplash-BE/internal/vo"
	"github.com/NganJason/Unsplash-BE/pkg/cerr"
	"github.com/NganJason/Unsplash-BE/pkg/server"
)

func SeedDataProcessor(
	ctx context.Context,
	writer http.ResponseWriter,
	req *http.Request,
) *server.HandlerResp {
	response := &vo.SeedDataResponse{}

	file := ctx.Value(internal.CtxFormBodyImg)
	if file == nil {
		return server.NewHandlerResp(
			response,
			cerr.New(
				"cannot parse file",
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

	profileImg := ctx.Value(internal.CtxProfileImg)
	if profileImg == nil {
		return server.NewHandlerResp(
			response,
			cerr.New(
				"cannot parse profileImg",
				http.StatusBadGateway,
			),
		)
	}

	profileImgBytes, _ := profileImg.([]byte)
	if profileImgBytes == nil {
		return server.NewHandlerResp(
			response,
			cerr.New(
				"assert profileImg error",
				http.StatusBadGateway,
			),
		)
	}

	val := ctx.Value(internal.CtxFormBodyVal).(string)
	var request vo.SeedDataRequest

	err := json.Unmarshal([]byte(val), &request)
	if err != nil {
		return server.NewHandlerResp(
			response,
			cerr.New(
				fmt.Sprintf("unmarshal json err=%s", err.Error()),
				http.StatusBadGateway,
			),
		)
	}

	p := &seedDataProcessor{
		ctx:             ctx,
		writer:          writer,
		fileBytes:       fileBytes,
		profileImgBytes: profileImgBytes,
		req:             &request,
		resp:            response,
	}

	return p.process()
}

type seedDataProcessor struct {
	ctx             context.Context
	writer          http.ResponseWriter
	fileBytes       []byte
	profileImgBytes []byte
	req             *vo.SeedDataRequest
	resp            *vo.SeedDataResponse
}

func (p *seedDataProcessor) process() *server.HandlerResp {
	userDM := model.NewUserDM(p.ctx)
	imageDM := model.NewImageDM(p.ctx)
	imageService, err := service.NewImageService(p.ctx)
	if err != nil {
		return server.NewHandlerResp(
			p.resp,
			err,
		)
	}

	userHandler := handler.NewUserHandler(
		p.ctx,
		userDM,
	)
	userHandler.SetImageService(imageService)

	imageHandler := handler.NewImageHandler(
		p.ctx,
		imageDM,
	)
	imageHandler.SetImageService(imageService)
	imageHandler.SetUserDM(userDM)

	user, _ := userHandler.GetUser(
		nil,
		p.req.EmailAddress,
		p.req.Password,
	)

	if user == nil {
		var err error

		user, err = userHandler.CreateUser(
			&vo.CreateUserRequest{
				EmailAddress: p.req.EmailAddress,
				Password:     p.req.Password,
				Username:     p.req.Username,
				FirstName:    p.req.FirstName,
				LastName:     p.req.LastName,
			},
		)
		if err != nil {
			return server.NewHandlerResp(
				p.resp,
				err,
			)
		}

		user, err = userHandler.UpdateProfileImg(
			*user.ID,
			p.profileImgBytes,
		)
		if err != nil {
			return server.NewHandlerResp(
				p.resp,
				err,
			)
		}
	}

	image, err := imageHandler.UploadImage(
		p.fileBytes,
		*user.ID,
		p.req.ImageDesc,
	)
	if err != nil {
		return server.NewHandlerResp(
			p.resp,
			err,
		)
	}

	p.resp.Image = image

	return server.NewHandlerResp(
		p.resp,
		nil,
	)
}
