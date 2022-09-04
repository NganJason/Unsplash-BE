package model

type UserLikeDM interface {
	CreateUserLike(userID uint64, imageID uint64) error
}
