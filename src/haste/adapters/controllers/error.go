package controllers

import (
	errorutils "haste/pkg/utils/error"
	requestresponseutils "haste/pkg/utils/request_response"

	"github.com/beego/beego/v2/server/web"
)

type ErrorController struct {
	web.Controller
}

func (ec *ErrorController) Error404() {
	err := errorutils.NewError("Page Not Found.")
	ec.Data["json"] = requestresponseutils.PrepareResponse("ASDF", err, "Resource not found!")
	ec.ServeJSON()
}
