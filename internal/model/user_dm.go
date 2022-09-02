package model

import "context"

type userDM struct {
	ctx context.Context
}

func NewUserDM(ctx context.Context) UserDM {
	return &userDM{
		ctx: ctx,
	}
}

func (dm *userDM) GetUserByID(userID *uint64) (*User, error) {
	return nil, nil
}

func (dm *userDM) GetUserByEmail(email *string) (*User, error) {
	return nil, nil
}

func (dm *userDM) CreateUser(req *CreateUserReq) (*User, error) {
	return nil, nil
}
