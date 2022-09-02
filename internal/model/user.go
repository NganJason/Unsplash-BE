package model

type UserDM interface {
	GetUserByID(userID *uint64) (*User, error)
	GetUserByEmail(email *string) (*User, error)
}
