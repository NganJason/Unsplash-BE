package model

import "context"

type imageDM struct {
	ctx context.Context
}

func NewImageDM(ctx context.Context) ImageDM {
	return &imageDM{}
}

func (dm *imageDM) GetImages(cursor *uint64, pageSize uint32) ([]*Image, error) {
	return nil, nil
}
