package util

import (
	"context"
	"fmt"
	"strconv"

	"github.com/NganJason/Unsplash-BE/pkg/cookies"
)

func GetUserIDFromCookies(ctx context.Context) (*uint64, error) {
	cookieVal := cookies.GetClientCookieValFromCtx(ctx)
	if cookieVal == nil {
		return nil, fmt.Errorf("cookies not found")
	}

	userID, err := strconv.ParseUint(*cookieVal, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("parse cookieVal err=%s", err.Error())
	}

	return Uint64Ptr(userID), nil
}
