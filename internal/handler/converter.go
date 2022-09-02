package handler

import (
	"github.com/NganJason/BE-template/internal/model"
	"github.com/NganJason/BE-template/internal/vo"
)

func toVoUser(dbUser *model.User) *vo.User {
	return &vo.User{
		ID:           dbUser.ID,
		Username:     dbUser.Username,
		EmailAddress: dbUser.EmailAddress,
		FirstName:    dbUser.FirstName,
		LastName:     dbUser.LastName,
	}
}
