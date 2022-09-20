package vo

const (
	CmdHealthCheck      = "CmdHealthCheck"
	CmdVerifyUser       = "CmdVerifyUser"
	CmdCreateUser       = "CmdCreateUser"
	CmdGetUser          = "CmdGetUser"
	CmdGetImages        = "CmdGetImages"
	CmdUploadImage      = "CmdUploadImg"
	CmdAddDeltaImage    = "CmdAddDeltaImage"
	CmdDownloadImg      = "CmdDownloadImg"
	CmdLogin            = "CmdLogin"
	CmdLogout           = "CmdLogout"
	CmdSeedData         = "CmdSeedData"
	CmdUpdateProfileImg = "CmdUpdateProfileImg"
)

const (
	PathHealthCheck      = "/api/healthcheck"
	PathVerifyUser       = "/api/user/verify"
	PathCreateUser       = "/api/user/create"
	PathGetUser          = "/api/user/get"
	PathLogin            = "/api/user/login"
	PathLogout           = "/api/user/logout"
	PathGetImages        = "/api/image/get_all"
	PathUploadImage      = "/api/image/upload"
	PathAddDeltaImage    = "/api/image/add_delta"
	PathSeedData         = "/api/data/seed"
	PathUpdateProfileImg = "/api/user/profile"
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

type VerifyUserRequest struct{}

type VerifyUserResponse struct {
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

type GetUserRequest struct {
	UserID *uint64 `json:"user_id"`
}

type GetUserResponse struct {
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

type AddDeltaImageResponse struct {
	CommonResponse
	Image *Image `json:"image"`
}

type DeltaImage struct {
	Likes     *uint32 `json:"likes"`
	Downloads *uint32 `json:"downloads"`
}

type User struct {
	ID           *uint64 `json:"id"`
	Username     *string `json:"username"`
	EmailAddress *string `json:"email_address"`
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
	ProfileUrl   *string `json:"profile_url"`
}

type Image struct {
	ID        *uint64 `json:"id"`
	User      *User   `json:"user"`
	Url       *string `json:"url"`
	Desc      *string `json:"desc"`
	Likes     *uint32 `json:"likes"`
	Downloads *uint32 `json:"downloads"`
}

type UpdateProfileImgRequest struct{}

type UpdateProfileImgResponse struct {
	CommonResponse
	User *User `json:"user"`
}

type SeedDataRequest struct {
	EmailAddress *string `json:"email_address"`
	Password     *string `json:"password"`
	Username     *string `json:"username"`
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
	ImageDesc    *string `json:"image_desc"`
}

type SeedDataResponse struct {
	CommonResponse
	Image *Image `json:"image"`
}
