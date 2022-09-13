package main

import (
	"context"
	"fmt"

	"github.com/hbagdi/go-unsplash/unsplash"
	"golang.org/x/oauth2"
)

const accessToken = "fmKnZB45lSW5zZE2RtzBjZCTK8mJvGEhzbK_Cy-_wPA"

type photoDM struct {
	ctx    context.Context
	client *unsplash.Unsplash
}

func NewPhotoDM(ctx context.Context) *photoDM {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: fmt.Sprintf("Client-ID %s", accessToken),
		},
	)

	client := oauth2.NewClient(oauth2.NoContext, ts)

	unsplash := unsplash.New(client)

	return &photoDM{
		ctx:    ctx,
		client: unsplash,
	}
}

func (dm *photoDM) GetPhotos(
	page,
	perPage int,
) (
	*[]unsplash.Photo,
	*unsplash.Response,
	error,
) {
	opt := new(unsplash.ListOpt)
	opt.Page = page
	opt.PerPage = perPage

	photos, resp, err := dm.client.Photos.All(opt)
	if err != nil {
		return nil, nil, err
	}

	return photos, resp, nil
}
