package model

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/NganJason/Unsplash-BE/internal/config"
	"github.com/NganJason/Unsplash-BE/internal/model/query"
	"github.com/NganJason/Unsplash-BE/internal/util"
	"github.com/NganJason/Unsplash-BE/pkg/cerr"
)

type userLikeDM struct {
	ctx context.Context
	db  *sql.DB
}

func NewUserLikeDM(ctx context.Context) UserLikeDM {
	return &userLikeDM{
		ctx: ctx,
		db:  config.GetDBs().UnsplashDB,
	}
}

func (dm *userLikeDM) GetUserLikes(
	userID *uint64,
	imageID *uint64,
	cursor *uint64,
	pageSize *uint32,
) ([]*UserLike, error) {
	q := query.NewUserLikeQuery()

	if userID != nil {
		q.UserID(*userID)
	}

	if imageID != nil {
		q.ImageID(*imageID)
	}

	q.Cursor(cursor).
		PageSize(pageSize).
		OrderBy(util.StrPtr("created_at DESC"))

	baseQuery := fmt.Sprintf(
		`SELECT * FROM %s WHERE `,
		dm.getTableName(),
	)

	wheres, args := q.Build()

	rows, err := dm.db.Query(
		baseQuery+wheres,
		args...,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, cerr.New(
			fmt.Sprintf("query userLikes err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	var userLikes []*UserLike
	for rows.Next() {
		var userLike UserLike

		if err := rows.Scan(
			&userLike.ID,
			&userLike.UserID,
			&userLike.ImageID,
			&userLike.CreatedAt,
			&userLike.UpdatedAt,
		); err != nil {
			if err == sql.ErrNoRows {
				return userLikes, nil
			}

			return nil, cerr.New(
				fmt.Sprintf("query users from db err=%s", err.Error()),
				http.StatusBadGateway,
			)
		}
		userLikes = append(userLikes, &userLike)
	}

	return userLikes, nil
}
func (dm *userLikeDM) CreateUserLike(userID uint64, imageID uint64) error {
	q := fmt.Sprintf(
		`
		INSERT INTO %s 
		(user_id, image_id, created_at, updated_at) 
		VALUES(?, ?, ?, ?)
		`, dm.getTableName(),
	)

	_, err := dm.db.Exec(
		q,
		userID,
		imageID,
		time.Now().UTC().UnixNano(),
		time.Now().UTC().UnixNano(),
	)
	if err != nil {
		return cerr.New(
			fmt.Sprintf("insert user into db err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	return nil
}

func (dm *userLikeDM) getTableName() string {
	return "user_like_tab"
}
