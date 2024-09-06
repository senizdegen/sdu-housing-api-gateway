package user_service

type User struct {
	UUID        string `json:"uuid"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"-"`
	Role        string `json:"role"`
	FullName    string `json:"full_name"`
	AvatarURL   string `json:"avatar_url"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	JWTToken    string `json:"jwt"`
}

type CreateUserDTO struct {
	FullName       string `json:"full_name"`
	PhoneNumber    string `json:"phone_number"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

type SigninUserDTO struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type CreateUserResponse struct {
	UUID     string `json:"uuid"`
	JWTToken string `json:"jwt"`
}
