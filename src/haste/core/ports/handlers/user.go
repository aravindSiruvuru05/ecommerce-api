package handlers

type UserPort interface {
	GetAllUsers() ([]*UserResponse, error)
	CreateUsers(userForm CreateUserForm) (UserResponse, error)
}

type CreateUserForm struct {
	Username string `json:"username" valid:"Required"`
	Email    string `json:"email" valid:"Required"`
	Password string `json:"password" valid:"Required"`
	Fullname string `json:"fullname" valid:"Required"`
}

type UserResponse struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}
