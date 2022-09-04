package model

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/NganJason/BE-template/pkg/cerr"
)

type userLikeDM struct {
	ctx context.Context
	db  *sql.DB
}

func NewUserLikeDM(ctx context.Context) UserLikeDM {
	return &userLikeDM{
		ctx: ctx,
	}
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
