package service

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/NganJason/Unsplash-BE/pkg/cerr"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type ImageService interface {
	UploadImage(fileBytes []byte) (url string, err error)
}

type cloudinaryService struct {
	ctx context.Context
	cloud *cloudinary.Cloudinary
}

const (
	cloudName = "di5bqamzn"
	apiKey = "177489556145968"
	apiSecret = "JaL7x6FqJH9QHx-mfuZI_IbIJX8"
	folder = "test_img"
)

func NewImageService(ctx context.Context) (ImageService, error) {
	cloudinaryClient, err := cloudinary.NewFromParams(
		cloudName,
		apiKey,
		apiSecret,
	)
	if err != nil {
		return nil, err
	}

	return &cloudinaryService{
		ctx: ctx,
		cloud: cloudinaryClient,
	}, nil
}

func (s *cloudinaryService) UploadImage(fileBytes []byte) (url string, err error) {
	reader := bytes.NewReader(fileBytes)

	resp, err := s.cloud.Upload.Upload(
		s.ctx,
		reader,
		uploader.UploadParams{
			Folder: folder,
		},
	)
	if err != nil {
		return "", cerr.New(
			fmt.Sprintf("upload image to cloudinary err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	return resp.URL, nil
}
