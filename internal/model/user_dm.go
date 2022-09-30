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

func (dm *userDM) GetUser(userID *uint64, username *string) (*User, error) {
	if userID == nil && username == nil {
		return nil, cerr.New(
			"userID and username cannot both be nil",
			http.StatusBadRequest,
		)
	}

	q := query.NewUserQuery()
	if username != nil {
		q.Username(*username)
	}

	if userID != nil {
		q.ID(*userID)
	}

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

	var users []*User
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
			&user.ProfileUrl,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			if err == sql.ErrNoRows {
				return nil, cerr.New(
					"user not found",
					http.StatusBadRequest,
				)
			}

			return nil, cerr.New(
				fmt.Sprintf("query users from db err=%s", err.Error()),
				http.StatusBadGateway,
			)
		}

		users = append(users, &user)
	}

	return users[0], nil
}

func (dm *userDM) GetUserByIDs(userIDs []uint64) ([]*User, error) {
	if len(userIDs) == 0 {
		return []*User{}, nil
	}

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
			&user.ProfileUrl,
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
			&user.ProfileUrl,
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

func (dm *userDM) UpdateUser(req *UpdateUserReq) (*User, error) {
	if req.UserID == 0 {
		return nil, cerr.New(
			"userID cannot be empty for update",
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

	q := query.NewUserQuery().ID(req.UserID)
	wheres, args := q.Build()
	finalQuery := baseQuery + wheres + "FOR UPDATE"

	var existingUser User
	err = tx.QueryRowContext(
		dm.ctx,
		finalQuery,
		args...,
	).Scan(
		&existingUser.ID,
		&existingUser.Username,
		&existingUser.EmailAddress,
		&existingUser.HashedPassword,
		&existingUser.Salt,
		&existingUser.LastName,
		&existingUser.FirstName,
		&existingUser.ProfileUrl,
		&existingUser.CreatedAt,
		&existingUser.UpdatedAt,
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("get existing user err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	if existingUser.ID == nil {
		return nil, cerr.New(
			"user does not exist for update",
			http.StatusBadRequest,
		)
	}

	if req.Username != nil {
		*existingUser.Username = *req.Username
	}

	if req.EmailAddress != nil {
		*existingUser.EmailAddress = *req.EmailAddress
	}

	if req.HashedPassword != nil {
		*existingUser.HashedPassword = *req.HashedPassword
	}

	if req.Salt != nil {
		*existingUser.Salt = *req.Salt
	}

	if req.LastName != nil {
		*existingUser.LastName = *req.LastName
	}

	if req.FirstName != nil {
		*existingUser.FirstName = *req.FirstName
	}

	if req.ProfileUrl != nil {
		existingUser.ProfileUrl = req.ProfileUrl
	}

	existingUser.UpdatedAt = util.Uint64Ptr(uint64(time.Now().UTC().UnixNano()))

	updateQuery := fmt.Sprintf(
		`
		UPDATE %s
		SET username = ?, email_address = ?, hashed_password = ?, salt = ?, last_name = ?, first_name = ?, profile_url = ?, updated_at = ?
		WHERE id = ?
		`,
		dm.getTableName(),
	)

	_, err = tx.ExecContext(
		dm.ctx,
		updateQuery,
		existingUser.Username,
		existingUser.EmailAddress,
		existingUser.HashedPassword,
		existingUser.Salt,
		existingUser.LastName,
		existingUser.FirstName,
		existingUser.ProfileUrl,
		existingUser.UpdatedAt,
		existingUser.ID,
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("update user err=%s", err.Error()),
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

	return &existingUser, nil
}

func (dm *userDM) SearchUsers(
	keyword string,
) (
	users []*User,
	err error,
) {
	baseQuery := fmt.Sprintf(
		`SELECT * from %s WHERE `,
		dm.getTableName(),
	)

	q := query.NewUserQuery().Keyword(&keyword)
	wheres, args := q.Build()

	rows, err := dm.db.Query(
		baseQuery+wheres,
		args...,
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("query users from db err=%s", err.Error()),
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
			&user.ProfileUrl,
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

func (dm *userDM) getTableName() string {
	return "user_tab"
}
