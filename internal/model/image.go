package model

type ImageDM interface {
	GetImages(cursor *uint64, pageSize uint32) ([]*Image, error)
	CreateImage(url string, userID uint64, desc *string) (*Image, error)
	AddDeltaImage(req *AddDeltaImageReq) (*Image, error)
}

type AddDeltaImageReq struct {
	ImageID   uint64
	Likes     *uint32
	Downloads *uint32
}
