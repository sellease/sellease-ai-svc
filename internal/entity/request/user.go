package request

type UserLoginRequest struct {
	Passcode string `json:"passcode" validate:"required"`
}

type AddUserImageRequest struct {
	Passcode  string `json:"passcode" validate:"required"`
	ImageName string `json:"image_name" validate:"required"`
}
