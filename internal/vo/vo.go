package vo

const (
	CmdHealthCheck   = "CmdHealthCheck"
	CmdGetUser       = "CmdGetUser"
	CmdCreateUser    = "CmdCreateUser"
	CmdGetImages     = "CmdGetImages"
	CmdUploadImage   = "CmdUploadImg"
	CmdAddDeltaImage = "CmdAddDeltaImage"
	CmdDownloadImg   = "CmdDownloadImg"
	CmdLogin         = "CmdLogin"
	CmdLogout        = "CmdLogout"
)

const (
	PathHealthCheck   = "/api/healthcheck"
	PathGetUser       = "/api/user/get"
	PathCreateUser    = "/api/user/create"
	PathLogin         = "/api/user/login"
	PathLogout        = "/api/user/logout"
	PathGetImages     = "/api/image/get_all"
	PathUploadImage   = "/api/image/upload"
	PathAddDeltaImage = "/api/image/add_delta"
)

type CommonResponse struct {
	DebugMsg *string `json:"debug_msg"`
}

type HealthCheckRequest struct {
	Message string `json:"message"`
}

type HealthCheckResponse struct {
	CommonResponse
	Message string `json:"message"`
}

type GetUserRequest struct{}

type GetUserResponse struct {
	CommonResponse
	User *User `json:"user"`
}

type LoginRequest struct {
	EmailAddress *string `json:"email_address"`
	Password     *string `json:"password"`
}

type LoginResponse struct {
	CommonResponse
	User *User `json:"user"`
}

type CreateUserRequest struct {
	EmailAddress *string `json:"email_address"`
	Password     *string `json:"password"`
	Username     *string `json:"username"`
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
}

type CreateUserResponse struct {
	CommonResponse
	User *User `json:"user"`
}

type GetImagesRequest struct {
	PageSize *uint32 `json:"page_size"`
	Cursor   *string `json:"cursor"`
}

type GetImagesResponse struct {
	CommonResponse
	Images     []*Image `json:"images"`
	NextCursor *string  `json:"next_cursor"`
}

type UploadImageRequest struct {
	Desc *string `json:"desc"`
}

type UploadImageResponse struct {
	CommonResponse
	Image *Image `json:"image"`
}

type AddDeltaImageRequest struct {
	ImageID    *uint64    `json:"image_id"`
	DeltaImage DeltaImage `json:"delta_image"`
}

type DeltaImage struct {
	Likes     *uint32 `json:"likes"`
	Downloads *uint32 `json:"downloads"`
}

type AddDeltaImageResponse struct {
	CommonResponse
	Image *Image `json:"image"`
}

type User struct {
	ID           *uint64 `json:"id"`
	Username     *string `json:"username"`
	EmailAddress *string `json:"email_address"`
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
}

type Image struct {
	ID        *uint64 `json:"id"`
	User      *User   `json:"user"`
	Url       *string `json:"url"`
	Desc      *string `json:"desc"`
	Likes     *uint32 `json:"likes"`
	Downloads *uint32 `json:"downloads"`
}
