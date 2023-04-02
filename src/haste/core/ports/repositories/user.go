package repositories

import (
	"context"
	db "haste/infra/driven/database/sqlc"
	"time"
)

type UserPort interface {
	GetAllUsersInDB(ctx context.Context) ([]User, error)
	CreatelUsersInDB(ctx context.Context, arg db.CreateUserParams) (User, error)
}

type User struct {
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}
