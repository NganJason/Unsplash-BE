package handler

import (
	"context"
	b64 "encoding/base64"
	"encoding/binary"
	"fmt"
	"net/http"
	"time"

	"github.com/NganJason/BE-template/internal/model"
	"github.com/NganJason/BE-template/internal/service"
	"github.com/NganJason/BE-template/internal/util"
	"github.com/NganJason/BE-template/internal/vo"
	"github.com/NganJason/BE-template/pkg/cerr"
)

type imageHandler struct {
	ctx          context.Context
	imageDM      model.ImageDM
	userDM       model.UserDM
	imageService service.ImageService
}

func NewImageHandler(
	ctx context.Context,
	imageDM model.ImageDM,
) *imageHandler {
	return &imageHandler{
		ctx:     ctx,
		imageDM: imageDM,
	}
}

func (h *imageHandler) SetUserDM(
	userDM model.UserDM,
) {
	h.userDM = userDM
}

func (h *imageHandler) SetImageService(
	imageService service.ImageService,
) {
	h.imageService = imageService
}

func (h *imageHandler) GetImages(
	pageSize *uint32,
	cursor *string,
) ([]*vo.Image, *string, error) {
	counter := *pageSize + 1

	cursorTimestamp, err := h.getCursorTimestamp(cursor)
	if err != nil {
		return nil, nil, err
	}

	images, err := h.imageDM.GetImages(
		util.Uint64Ptr(cursorTimestamp),
		counter,
	)
	if err != nil {
		return nil, nil, err
	}

	nextCursor := h.getNextCursor(images, counter)

	if len(images) == int(counter) {
		images = images[:len(images)-1]
	}

	userIDs := h.extractUserIDs(images)
	users, err := h.userDM.GetUserByIDs(userIDs)
	if err != nil {
		return nil, nil, err
	}
	userIDMap := h.getUserIDMap(users)

	return toVoImages(images, userIDMap), nextCursor, nil
}

func (h *imageHandler) UploadImage(
	fileBytes []byte,
	userID uint64,
	desc *string,
) (*vo.Image, error) {
	url, err := h.imageService.UploadImage(fileBytes)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("upload img err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	image, err := h.imageDM.CreateImage(url, userID, desc)

	userIDs := h.extractUserIDs([]*model.Image{image})
	users, err := h.userDM.GetUserByIDs(userIDs)
	if err != nil {
		return nil, err
	}
	userIDMap := h.getUserIDMap(users)

	return toVoImage(image, userIDMap), nil
}

func (h *imageHandler) getCursorTimestamp(cursor *string) (uint64, error) {
	if cursor == nil {
		return uint64(time.Now().Unix()), nil
	}

	cursorByte, err := b64.StdEncoding.DecodeString(*cursor)
	if err != nil {
		return 0, cerr.New(
			fmt.Sprintf("decode cursor err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	cursorTimestamp := uint64(binary.LittleEndian.Uint64(cursorByte))

	return cursorTimestamp, nil
}

func (h *imageHandler) getNextCursor(images []*model.Image, counter uint32) *string {
	var nextCursor *string

	if len(images) == int(counter) {
		nextCursorTimestamp := images[len(images)-1].CreatedAt
		cursorByte := make([]byte, 8)
		binary.LittleEndian.PutUint64(cursorByte, *nextCursorTimestamp)

		nextCursor = util.StrPtr(b64.StdEncoding.EncodeToString(cursorByte))
	}

	return nextCursor
}

func (h *imageHandler) getUserIDMap(users []*model.User) map[uint64]*model.User {
	userIDMap := make(map[uint64]*model.User)

	for _, user := range users {
		userIDMap[*user.ID] = user
	}

	return userIDMap
}

func (h *imageHandler) extractUserIDs(images []*model.Image) []uint64 {
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
