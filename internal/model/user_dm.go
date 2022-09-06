package model

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/NganJason/Unsplash-BE/internal/config"
	"github.com/NganJason/Unsplash-BE/internal/model/query"
	"github.com/NganJason/Unsplash-BE/pkg/cerr"
)

type userDM struct {
	ctx context.Context
	db  *sql.DB
}

func NewUserDM(ctx context.Context) UserDM {
	return &userDM{
		ctx: ctx,
		db:  config.GetDBs().UnsplashDB,
	}
}

func (dm *userDM) GetUserByIDs(userIDs []uint64) ([]*User, error) {
	var users []*User

	q := query.NewUserQuery()
	q.IDs(userIDs)

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
			fmt.Sprintf("query users err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	for rows.Next() {
		var user User

		if err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.EmailAddress,
			&user.HashedPassword,
			&user.Salt,
			&user.LastName,
			&user.FirstName,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			if err == sql.ErrNoRows {
				return users, nil
			}

			return nil, cerr.New(
				fmt.Sprintf("query users from db err=%s", err.Error()),
				http.StatusBadGateway,
			)
		}
		users = append(users, &user)
	}

	return users, nil
}

func (dm *userDM) GetUserByEmails(emails []string) ([]*User, error) {
	var users []*User

	q := query.NewUserQuery()
	q.Emails(emails)

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
			fmt.Sprintf("query users err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	for rows.Next() {
		var user User

		if err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.EmailAddress,
			&user.HashedPassword,
			&user.Salt,
			&user.LastName,
			&user.FirstName,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			if err == sql.ErrNoRows {
				return users, nil
			}

			return nil, cerr.New(
				fmt.Sprintf("query users from db err=%s", err.Error()),
				http.StatusBadGateway,
			)
		}
		users = append(users, &user)
	}

	return users, nil
}

func (dm *userDM) CreateUser(req *CreateUserReq) (*User, error) {
	q := fmt.Sprintf(
		`
		INSERT INTO %s 
		(username, email_address, hashed_password, salt, last_name, first_name, created_at, updated_at) 
		VALUES(?, ?, ?, ?, ?, ?, ?, ?)
		`, dm.getTableName(),
	)

	result, err := dm.db.Exec(
		q,
		req.Username,
		req.EmailAddress,
		req.HashedPassword,
		req.SaltString,
		req.LastName,
		req.FirstName,
		time.Now().UTC().UnixNano(),
		time.Now().UTC().UnixNano(),
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("insert user into db err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	lastInsertID, _ := result.LastInsertId()

	users, err := dm.GetUserByIDs(
		[]uint64{uint64(lastInsertID)},
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("refetch user from db err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	return users[0], nil
}

func (dm *userDM) getTableName() string {
	return "user_tab"
}
