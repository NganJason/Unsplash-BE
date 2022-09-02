package model

type UserDM interface {
	GetUserByID(userID *uint64) (*User, error)
	GetUserByEmail(email *string) (*User, error)

	CreateUser(req *CreateUserReq) (*User, error)
}

type CreateUserReq struct {
	EmailAddress   *string
	HashedPassword *string
	SaltString     *string
	Username       *string
	FirstName      *string
	LastName       *string
}
