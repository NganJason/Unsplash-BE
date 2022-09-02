package model

type User struct {
	ID             *uint64 `json:"id"`
	Username       *string `json:"username"`
	EmailAddress   *string `json:"email_address"`
	HashedPassword *string `json:"hashed_password"`
	Salt           *string `json:"salt"`
	LastName       *string `json:"last_name"`
	FirstName      *string `json:"first_name"`
	CreatedAt      *uint64 `json:"created_at"`
	UpdatedAt      *uint64 `json:"updated_at"`
}

type Image struct {
	ID        *uint64 `json:"id"`
	UserID    *uint64 `json:"user_id"`
	Url       *string `json:"url"`
	Desc      *string `json:"desc"`
	Likes     *uint32 `json:"likes"`
	Downloads *uint32 `json:"downloads"`
	CreatedAt *uint64 `json:"created_at"`
	UpdatedAt *uint64 `json:"updated_at"`
}

type Tag struct {
	ID        *uint64 `json:"id"`
	Name      *string `json:"name"`
	CreatedAt *uint64 `json:"created_at"`
	UpdatedAt *uint64 `json:"updated_at"`
}

type ImageTag struct {
	ID        *uint64 `json:"id"`
	ImageID   *uint64 `json:"image_id"`
	CreatedAt *uint64 `json:"created_at"`
	UpdatedAt *uint64 `json:"updated_at"`
}

type UserLike struct {
	UserID    *uint64 `json:"user_id"`
	ImageID   *uint64 `json:"image_id"`
	CreatedAt *uint64 `json:"created_at"`
	UpdatedAt *uint64 `json:"updated_at"`
}
