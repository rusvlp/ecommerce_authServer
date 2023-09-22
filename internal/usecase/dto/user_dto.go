package dto

type CreateUserDTO struct {
	Username             string
	Password             string
	PasswordConfirmation string
	Email                string
}

func (u *CreateUserDTO) String() string {
	return "User " + u.Username
}

type GetUserDTO struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
