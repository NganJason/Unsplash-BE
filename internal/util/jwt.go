package util

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/NganJason/Unsplash-BE/pkg/auth"
	"github.com/NganJason/Unsplash-BE/pkg/cookies"
)

func GenerateTokenAndAddCookies(ctx context.Context, value string) error {
	jwt, err := GenerateJWTToken(value)
	if err != nil {
		return fmt.Errorf(
			"generate jwt token err=%s", err.Error(),
		)
	}

	c := cookies.CreateCookie(jwt)

	cookies.AddServerCookieToCtx(ctx, c)

	return nil
}

func GenerateCookies(value string) (*http.Cookie, error) {
	jwt, err := GenerateJWTToken(value)
	if err != nil {
		return nil, err
	}

	c := cookies.CreateCookie(jwt)

	return c, nil
}

func GenerateJWTToken(value string) (string, error) {
	secretKey, err := GetDotEnvVariable(JWTSecretEnvName)
	if err != nil {
		return "", err
	}

	expirationMinuteString, err := GetDotEnvVariable(JWTExpirationMinutesEnvName)
	if err != nil {
		return "", err
	}

	expirationMinute, err := strconv.Atoi(expirationMinuteString)
	if err != nil {
		return "", err
	}

	jwtToken, err := auth.GenerateJWTToken(value, secretKey, expirationMinute)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func ParseJWTToken(tokenString string) (*auth.Claims, error) {
	secretKey, err := GetDotEnvVariable(JWTSecretEnvName)
	if err != nil {
		return &auth.Claims{}, err
	}

	return auth.ParseJWTToken(tokenString, secretKey)
}
