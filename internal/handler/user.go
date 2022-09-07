package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/NganJason/Unsplash-BE/internal/model"
	"github.com/NganJason/Unsplash-BE/internal/util"
	"github.com/NganJason/Unsplash-BE/internal/vo"
	"github.com/NganJason/Unsplash-BE/pkg/auth"
	"github.com/NganJason/Unsplash-BE/pkg/cerr"
)

type userHandler struct {
	ctx        context.Context
	userDM     model.UserDM
	userLikeDM model.UserLikeDM
	imageDM    model.ImageDM
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

func (h *userHandler) likeImage(userID, imageID uint64) error {
	userLikes, err := h.userLikeDM.GetUserLikes(
		&userID,
		&imageID,
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
