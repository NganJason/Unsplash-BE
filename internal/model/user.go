package model

type UserDM interface {
	GetUserByIDs(userID []uint64) ([]*User, error)
	GetUserByEmails(email []string) ([]*User, error)

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
