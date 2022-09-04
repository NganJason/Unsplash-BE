package service

import "context"

type ImageService interface {
	UploadImage(fileBytes []byte) (url string, err error)
}

type imageService struct {
	ctx context.Context
}

func NewImageService(ctx context.Context) ImageService {
	return &imageService{
		ctx: ctx,
	}
}

func (s *imageService) UploadImage(fileBytes []byte) (url string, err error) {
	return "", nil
}
