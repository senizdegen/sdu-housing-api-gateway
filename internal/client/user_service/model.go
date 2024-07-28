package user_service

type User struct {
	UUID     string `json:"uuid"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
