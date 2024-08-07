package user_service

type User struct {
	UUID        string `json:"uuid"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"-"`
	Role        string `json:"role"`
	JWTToken    string `json:"jwt"`
}

type CreateUserDTO struct {
	PhoneNumber    string `json:"phone_number"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

type SigninUserDTO struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
