package model

type UserLikeDM interface {
	CreateUserLike(userID uint64, imageID uint64) error
	GetUserLikes(userID *uint64, imageID *uint64, cursor *uint64, pageSize *uint32) ([]*UserLike, error)
}
