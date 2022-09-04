package model

type ImageDM interface {
	GetImages(cursor *uint64, pageSize uint32) ([]*Image, error)
}
