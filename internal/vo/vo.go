package vo

const (
	CmdHealthCheck = "CmdHealthCheck"
	CmdGetUser     = "CmdGetUser"
	CmdCreateUser  = "CmdCreateUser"
	CmdGetImages   = "CmdGetImages"
	CmdUploadImage = "CmdUploadImg"
	CmdLikeImage   = "CmdLikeImage"
	CmdDownloadImg = "CmdDownloadImg"
)

const (
	PathHealthCheck = "/api/healthcheck"
	PathGetUser     = "/api/user/get"
	PathCreateUser  = "/api/user/create"
	PathGetImages   = "/api/image/get_all"
	PathUploadImage = "/api/image/upload"
	PathLikeImage   = "/api/image/like"
	PathDownloadImg = "/api/image/downlaod"
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

type GetUserRequest struct {
	EmailAddress *string `json:"email_address"`
	Password     *string `json:"password"`
}

type GetUserResponse struct {
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
	Cursor   *uint64 `json:"cursor"`
}

type GetImagesResponse struct {
	CommonResponse
	Images []*Image `json:"images"`
}

type UploadImageRequest struct{}

type UploadImageResponse struct {
	CommonResponse
	Image *Image `json:"image"`
}

type LikeImageRequest struct {
	ImageID *uint64 `json:"image_id"`
}

type LikeImageResponse struct {
	CommonResponse
}

type DownloadImageRequest struct {
	ImageID *uint64 `json:"image_id"`
}

type DownloadImageResponse struct {
	CommonResponse
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
	UserID    *uint64 `json:"user_id"`
	Url       *string `json:"url"`
	Desc      *string `json:"desc"`
	Likes     *uint32 `json:"likes"`
	Downloads *uint32 `json:"downloads"`
}
