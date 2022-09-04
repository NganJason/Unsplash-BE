package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/NganJason/BE-template/internal/model"
	"github.com/NganJason/BE-template/internal/vo"
	"github.com/NganJason/BE-template/pkg/auth"
	"github.com/NganJason/BE-template/pkg/cerr"
)

type userHandler struct {
	ctx    context.Context
	userDM model.UserDM
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
			fmt.Sprintf("get user by username err=%s", err.Error()),
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
	hashedPassword, saltString := auth.CreatePasswordSHA(*req.Password, 16)

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
