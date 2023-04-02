package user

import (
	"encoding/json"
	"fmt"
	"haste/adapters/controllers"
	"haste/core/ports/handlers"
	requestresponse "haste/pkg/utils/request_response"

	"github.com/beego/beego/v2/adapter/validation"
)

// Operations about Users
type UserController struct {
	controllers.BaseController
	Component handlers.UserPort
}

// UpdateComponent is used to instantiate a component inside the controller.
func (c *UserController) UpdateComponent(component interface{}) {
	c.Component, _ = component.(handlers.UserPort)
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (c *UserController) CreateUser() {
	var data handlers.UserResponse
	var err error

	var form handlers.CreateUserForm
	fmt.Println("form", form)
	if err = json.Unmarshal(c.GetRequestBody(), &form); err == nil {
		validation := validation.Validation{}
		if valid, _ := validation.Valid(&form); !valid {
			c.Error(err)
		} else if data, err = c.Component.CreateUsers(form); err == nil {
			c.Data["json"] = requestresponse.PrepareResponse(data, err, "Success")
			c.ServeJSON()
		}
	}
	c.Error(err)
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAllUsers() {
	data, err := u.Component.GetAllUsers()
	u.Data["json"] = requestresponse.PrepareResponse(data, err, "Success")
	_ = u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) GetById() {
	uid := u.Ctx.Input.Param(":id")
	u.Data["json"] = uid
	u.ServeJSON()
}

func init() {
	fmt.Println("user controller -----")
}
