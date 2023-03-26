package repositories

type UserPort interface {
	GetAllUsersInDB() []*User
}

type User struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
