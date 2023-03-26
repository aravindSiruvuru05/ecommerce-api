package user

import (
	"haste/core/components"
	"haste/core/ports/handlers"
	"haste/core/ports/repositories"

	maputils "haste/pkg/utils/map"
)

type UserComponent struct {
	components.BaseComponent
}

func (c *UserComponent) GetAllUsers() []*handlers.UserResponse {
	urepo := c.GetRepository("User").(repositories.UserPort)
	users := urepo.GetAllUsersInDB()
	var result []*handlers.UserResponse
	for _, user := range users {
		var ur handlers.UserResponse
		_ = maputils.MapInterfaceToObject(user, &ur)
		result = append(result, &ur)
	}
	return result
}

func init() {
	components.ComponentMap["User"] = func(bc *components.BaseComponent) interface{} {
		c := &UserComponent{
			BaseComponent: *bc,
		}
		return handlers.UserPort(c)
	}
}
