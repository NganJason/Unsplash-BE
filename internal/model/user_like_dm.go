package model

import "context"

type userLikeDM struct {
	ctx context.Context
}

func NewUserLikeDM(ctx context.Context) UserLikeDM {
	return &userLikeDM{
		ctx: ctx,
	}
}

func (dm *userLikeDM) CreateUserLike(userID uint64, imageID uint64) error {
	return nil
}
