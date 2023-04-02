package user

import (
	"haste/core/components"
	"haste/core/ports/handlers"
	"haste/core/ports/repositories"
	db "haste/infra/driven/database/sqlc"

	maputils "haste/pkg/utils/map"
	passwordutils "haste/pkg/utils/password"
)

type UserComponent struct {
	components.BaseComponent
}

func (c *UserComponent) CreateUsers(userForm handlers.CreateUserForm) (handlers.UserResponse, error) {
	var result handlers.UserResponse
	var err error

	urepo := c.GetRepository("User").(repositories.UserPort)

	hashedPassword, err := passwordutils.HashPassword(userForm.Password)
	if err != nil {
		return result, err
	}

	ur, err := urepo.CreatelUsersInDB(c.ReqCtx, db.CreateUserParams{
		Username:       userForm.Username,
		Email:          userForm.Email,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		return result, err
	}

	result = handlers.UserResponse{
		Username: ur.Username,
		Email:    ur.Email,
	}
	return result, err
}

func (c *UserComponent) GetAllUsers() ([]*handlers.UserResponse, error) {
	urepo := c.GetRepository("User").(repositories.UserPort)
	users, err := urepo.GetAllUsersInDB(c.ReqCtx)
	var result []*handlers.UserResponse
	for _, user := range users {
		var ur handlers.UserResponse
		_ = maputils.MapInterfaceToObject(user, &ur)
		result = append(result, &ur)
	}
	return result, err
}

func init() {
	components.ComponentMap["User"] = func(bc *components.BaseComponent) interface{} {
		c := &UserComponent{
			BaseComponent: *bc,
		}
		return handlers.UserPort(c)
	}
}
