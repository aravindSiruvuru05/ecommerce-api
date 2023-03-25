package user

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

// Operations about Users
type UserController struct {
	web.Controller
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
func (u *UserController) GetAll() {
	u.Data["json"] = "alhl"
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
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
