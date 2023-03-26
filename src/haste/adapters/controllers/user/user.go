package user

import (
	"fmt"
	"haste/adapters/controllers"
	"haste/core/ports/handlers"
	requestresponse "haste/pkg/utils/request_response"
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
func (u *UserController) Post() {
	u.Data["json"] = map[string]string{"uid": "uid"}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAllUsers() {
	result := u.Component.GetAllUsers()
	u.Data["json"] = requestresponse.PrepareResponse(result, nil, 200)
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
