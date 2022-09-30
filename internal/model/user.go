package model

type UserDM interface {
	GetUserByIDs(userID []uint64) ([]*User, error)
	GetUserByEmails(email []string) ([]*User, error)

	CreateUser(req *CreateUserReq) (*User, error)
	UpdateUser(req *UpdateUserReq) (*User, error)
	SearchUsers(keyword string) ([]*User, error)
}

type CreateUserReq struct {
	EmailAddress   *string
	HashedPassword *string
	SaltString     *string
	Username       *string
	FirstName      *string
	LastName       *string
}

type UpdateUserReq struct {
	UserID         uint64
	EmailAddress   *string
	HashedPassword *string
	Salt           *string
	Username       *string
	FirstName      *string
	LastName       *string
	ProfileUrl     *string
}
