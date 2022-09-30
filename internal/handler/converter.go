package handler

import (
	"github.com/NganJason/Unsplash-BE/internal/model"
	"github.com/NganJason/Unsplash-BE/internal/vo"
)

func toVoUser(dbUser *model.User) *vo.User {
	return &vo.User{
		ID:           dbUser.ID,
		Username:     dbUser.Username,
		EmailAddress: dbUser.EmailAddress,
		FirstName:    dbUser.FirstName,
		LastName:     dbUser.LastName,
		ProfileUrl:   dbUser.ProfileUrl,
	}
}

func toVoUsers(dbUsers []*model.User) []*vo.User {
	voUsers := make([]*vo.User, 0)

	for _, user := range dbUsers {
		voUsers = append(voUsers, toVoUser(user))
	}

	return voUsers
}

func toVoImages(dbImages []*model.Image, userIDMap map[uint64]*model.User) []*vo.Image {
	voImages := make([]*vo.Image, 0)

	for _, img := range dbImages {
		voImages = append(
			voImages,
			toVoImage(img, userIDMap),
		)
	}

	return voImages
}

func toVoImage(dbImage *model.Image, userIDMap map[uint64]*model.User) *vo.Image {
	var user *vo.User
	if dbUser, ok := userIDMap[*dbImage.UserID]; ok {
		user = toVoUser(dbUser)
	}

	return &vo.Image{
		ID:        dbImage.ID,
		User:      user,
		Url:       dbImage.Url,
		Desc:      dbImage.Desc,
		Likes:     dbImage.Likes,
		Downloads: dbImage.Downloads,
	}
}
