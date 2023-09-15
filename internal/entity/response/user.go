package response

type User struct {
	ID string `json:"id" gorm:"primaryKey"`
}

type UserLoginResponse struct {
	AccessToken string `json:"access_token"`
}
