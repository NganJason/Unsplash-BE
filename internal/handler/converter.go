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

func toVoImages(dbImages []*model.Image, userIDMap map[uint64]*model.User) []*vo.Image {
	voImages := make([]*vo.Image, 0)

	for _, img := range dbImages {
		voImages = append(
			voImages,
			&vo.Image{
				ID:        img.ID,
				User:      toVoUser(userIDMap[*img.UserID]),
				Url:       img.Url,
				Desc:      img.Desc,
				Likes:     img.Likes,
				Downloads: img.Downloads,
			},
		)
	}

	return voImages
}
