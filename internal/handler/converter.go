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
	}
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
	return &vo.Image{
		ID:        dbImage.ID,
		User:      toVoUser(userIDMap[*dbImage.UserID]),
		Url:       dbImage.Url,
		Desc:      dbImage.Desc,
		Likes:     dbImage.Likes,
		Downloads: dbImage.Downloads,
	}
}
