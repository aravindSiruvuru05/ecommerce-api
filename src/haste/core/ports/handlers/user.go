package handlers

type UserPort interface {
	GetAllUsers() []*UserResponse
}

type UserResponse struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type UserRequestForm struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
