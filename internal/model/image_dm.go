package model

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/NganJason/BE-template/internal/model/query"
	"github.com/NganJason/BE-template/internal/util"
	"github.com/NganJason/BE-template/pkg/cerr"
)

type imageDM struct {
	ctx context.Context
	db  *sql.DB
}

func NewImageDM(ctx context.Context) ImageDM {
	return &imageDM{}
}

func (dm *imageDM) GetImages(cursor *uint64, pageSize uint32) ([]*Image, error) {
	return nil, nil
}

func (dm *imageDM) GetImageByID(id uint64) (*Image, error) {
	q := query.NewImageQuery()
	q.ID(id)

	wheres, args := q.Build()
	baseQuery := fmt.Sprintf(
		`SELECT * from %s WHERE `,
		dm.getTableName(),
	)

	var image Image
	err := dm.db.QueryRow(
		baseQuery+wheres,
		args...,
	).Scan(
		&image.ID,
		&image.UserID,
		&image.Url,
		&image.Desc,
		&image.Likes,
		&image.Downloads,
		&image.CreatedAt,
		&image.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, cerr.New(
			fmt.Sprintf("query image err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	return &image, nil
}

func (dm *imageDM) CreateImage(
	url string,
	userID uint64,
	desc *string,
) (*Image, error) {
	q := fmt.Sprintf(
		`
		INSERT INTO %s 
		(user_id, url, desc, likes, downloads, created_at, updated_at) 
		VALUES(?, ?, ?, ?, ?, ?, ?)
		`, dm.getTableName(),
	)

	result, err := dm.db.Exec(
		q,
		userID,
		url,
		desc,
		0,
		0,
		time.Now().UTC().UnixNano(),
		time.Now().UTC().UnixNano(),
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("insert image into db err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	lastInsertID, _ := result.LastInsertId()

	image, err := dm.GetImageByID(
		uint64(lastInsertID),
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("refetch image from db err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	return image, nil
}

func (dm *imageDM) AddDeltaImage(req *AddDeltaImageReq) (*Image, error) {
	if req.ImageID == 0 {
		return nil, cerr.New(
			"imageID cannot be empty for update",
			http.StatusBadRequest,
		)
	}

	tx, err := dm.db.BeginTx(dm.ctx, nil)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("begin tx for update err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}
	defer tx.Rollback()

	baseQuery := fmt.Sprintf(
		`SELECT * from %s WHERE `,
		dm.getTableName(),
	)

	q := query.NewImageQuery().ID(req.ImageID)
	wheres, args := q.Build()
	finalQuery := baseQuery + wheres + "FOR UPDATE"

	var existingImage Image
	err = tx.QueryRowContext(
		dm.ctx,
		finalQuery,
		args...,
	).Scan(
		&existingImage.ID,
		&existingImage.UserID,
		&existingImage.Url,
		&existingImage.Desc,
		&existingImage.Likes,
		&existingImage.Downloads,
		&existingImage.CreatedAt,
		&existingImage.UpdatedAt,
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("get existing image err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	if existingImage.ID == nil {
		return nil, cerr.New(
			"image does not exist for update",
			http.StatusBadRequest,
		)
	}

	if req.Downloads != nil {
		existingImage.Downloads = req.Downloads
	}

	if req.Likes != nil {
		existingImage.Likes = req.Likes
	}

	existingImage.UpdatedAt = util.Uint64Ptr(uint64(time.Now().UTC().UnixNano()))

	updateQuery := fmt.Sprintf(
		`
		UPDATE %s
		SET likes = ?, downloads = ?
		WHERE id = ?
		`,
		dm.getTableName(),
	)

	_, err = tx.ExecContext(
		dm.ctx,
		updateQuery,
		existingImage.Likes,
		existingImage.Downloads,
		existingImage.ID,
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("update image err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	err = tx.Commit()
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("commit transaction err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	return &existingImage, nil
}

func (dm *imageDM) getTableName() string {
	return "image_tab"
}
