package model

type ImageDM interface {
	GetImages(cursor *uint64, pageSize uint32) ([]*Image, error)
	UploadImage(url string, userID uint64, desc *string) (*Image, error)
}
