package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/NganJason/Unsplash-BE/internal/model"
	"github.com/NganJason/Unsplash-BE/internal/service"
	"github.com/NganJason/Unsplash-BE/internal/util"
	"github.com/NganJason/Unsplash-BE/internal/vo"
	"github.com/NganJason/Unsplash-BE/pkg/auth"
	"github.com/NganJason/Unsplash-BE/pkg/cerr"
)

type userHandler struct {
	ctx          context.Context
	userDM       model.UserDM
	userLikeDM   model.UserLikeDM
	imageDM      model.ImageDM
	imageService service.ImageService
}

func NewUserHandler(
	ctx context.Context,
	userDM model.UserDM,
) *userHandler {
	return &userHandler{
		ctx:    ctx,
		userDM: userDM,
	}
}

func (h *userHandler) SetUserLikeDM(userLikeDM model.UserLikeDM) {
	h.userLikeDM = userLikeDM
}

func (h *userHandler) SetImageDM(imageDM model.ImageDM) {
	h.imageDM = imageDM
}

func (h *userHandler) SetImageService(imageService service.ImageService) {
	h.imageService = imageService
}

func (h *userHandler) GetUser(
	userID *uint64,
	emailAddress *string,
	password *string,
) (*vo.User, error) {
	if userID != nil {
		users, err := h.userDM.GetUserByIDs([]uint64{*userID})
		if err != nil {
			return nil, cerr.New(
				fmt.Sprintf("get user by ID err=%s", err.Error()),
				http.StatusBadGateway,
			)
		}

		if len(users) == 0 {
			return nil, cerr.New(
				fmt.Sprintf("cannot find user with userID=%d", *userID),
				http.StatusBadRequest,
			)
		}

		return toVoUser(users[0]), nil
	}

	if emailAddress == nil || password == nil {
		return nil, cerr.New(
			fmt.Sprintf("email and password cannot be nil"),
			http.StatusBadRequest,
		)
	}

	users, err := h.userDM.GetUserByEmails([]string{*emailAddress})
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("get user by email err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	if len(users) == 0 {
		return nil, cerr.New(
			fmt.Sprintf("cannot find user with email=%s", *emailAddress),
			http.StatusBadGateway,
		)
	}

	user := users[0]
	isPasswordMatch := auth.ComparePasswordSHA(
		*password,
		*user.HashedPassword,
		*user.Salt,
	)

	if !isPasswordMatch {
		return nil, cerr.New(
			"invalid password",
			http.StatusForbidden,
		)
	}

	return toVoUser(user), nil
}

func (h *userHandler) CreateUser(req *vo.CreateUserRequest) (*vo.User, error) {
	hashedPassword, saltString := auth.CreatePasswordSHA(*req.Password, util.SaltSize)

	user, err := h.userDM.CreateUser(
		&model.CreateUserReq{
			EmailAddress:   req.EmailAddress,
			HashedPassword: &hashedPassword,
			SaltString:     &saltString,
			Username:       req.Username,
			FirstName:      req.FirstName,
			LastName:       req.LastName,
		},
	)
	if err != nil {
		return nil, err
	}

	return toVoUser(user), nil
}

func (h *userHandler) AddDeltaImage(userID uint64, imageID uint64, deltaImage *vo.DeltaImage) error {
	if deltaImage.Downloads != nil {
		h.downloadImage(userID, imageID)
	}

	if deltaImage.Likes != nil {
		err := h.likeImage(userID, imageID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *userHandler) UpdateProfileImg(userID uint64, fileBytes []byte) (*vo.User, error) {
	url, err := h.imageService.UploadImage(fileBytes)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("upload img err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	user, err := h.userDM.UpdateUser(&model.UpdateUserReq{
		UserID:     userID,
		ProfileUrl: util.StrPtr(url),
	})
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("update user err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	return toVoUser(user), nil
}

func (h *userHandler) GetUserLikes(userID uint64) ([]*vo.Image, error) {
	userLikes, err := h.userLikeDM.GetUserLikes(&userID, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	
	imageIDs := make([]uint64, 0)
	for _, userLike := range userLikes {
		imageIDs = append(imageIDs, *userLike.ImageID)
	}

	images, err := h.imageDM.GetImagesByIDs(imageIDs)
	if err != nil {
		return nil, err
	}

	userIDs := h.extractUserIDs(images)
	users, err := h.userDM.GetUserByIDs(userIDs)
	if err != nil {
		return nil, err
	}
	userIDMap := h.getUserIDMap(users)

	return toVoImages(images, userIDMap), nil
}

func (h *userHandler) likeImage(userID, imageID uint64) error {
	userLikes, err := h.userLikeDM.GetUserLikes(
		&userID,
		&imageID,
		nil,
		nil,
	)
	if err != nil {
		return err
	}

	if len(userLikes) > 0 {
		return cerr.New(
			"duplicate like, ignore",
			http.StatusBadRequest,
		)
	}

	err = h.userLikeDM.CreateUserLike(userID, imageID)
	if err != nil {
		return err
	}

	_, err = h.imageDM.AddDeltaImage(
		&model.AddDeltaImageReq{
			ImageID: imageID,
			Likes:   util.Uint32Ptr(uint32(1)),
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (h *userHandler) downloadImage(userID, imageID uint64) error {
	_, err := h.imageDM.AddDeltaImage(
		&model.AddDeltaImageReq{
			ImageID:   imageID,
			Downloads: util.Uint32Ptr(uint32(1)),
		},
	)
	fmt.Println("here", err)
	if err != nil {
		return err
	}

	return nil
}

func (h *userHandler) getUserIDMap(users []*model.User) map[uint64]*model.User {
	userIDMap := make(map[uint64]*model.User)

	for _, user := range users {
		userIDMap[*user.ID] = user
	}

	return userIDMap
}

func (h *userHandler) extractUserIDs(images []*model.Image) []uint64 {
	userIDs := make([]uint64, 0)
	userIDMap := make(map[uint64]bool)

	for _, img := range images {
		userIDMap[*img.UserID] = true
	}

	for userID := range userIDMap {
		userIDs = append(userIDs, userID)
	}

	return userIDs
}

