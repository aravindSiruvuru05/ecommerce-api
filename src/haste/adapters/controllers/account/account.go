package account

import (
	"encoding/json"
	"haste/adapters/controllers"
	"haste/core/ports/handlers"
	requestresponseutils "haste/pkg/utils/request_response"

	"github.com/beego/beego/v2/core/validation"
)

// Operations about Users
type AccountController struct {
	controllers.BaseController
	Component handlers.AccountPort
}

// UpdateComponent is used to instantiate a component inside the controller.
func (c *AccountController) UpdateComponent(component interface{}) {
	c.Component, _ = component.(handlers.AccountPort)
}

func (c *AccountController) CreateAccount() {
	var data handlers.AccountResponse
	var err error

	var form handlers.AccountForm
	if err = json.Unmarshal(c.GetRequestBody(), &form); err == nil {
		validation := validation.Validation{}
		if valid, _ := validation.Valid(&form); !valid {
			c.Error(err)
		} else if data, err = c.Component.CreateAccount(form); err == nil {
			c.Data["json"] = requestresponseutils.PrepareResponse(data, err, "Success")
			c.ServeJSON()
		}
	}
	c.Error(err)
}

func (c *AccountController) GetAccountById() {
	var data handlers.AccountResponse
	var err error

	if data, err = c.Component.GetAccountById(); err == nil {
		c.Data["json"] = requestresponseutils.PrepareResponse(data, err, "Success")
		c.ServeJSON()
	}
}

func (c *AccountController) TransferAmount() {
	var data handlers.TransferResponse
	var err error

	var form handlers.TransferForm
	if err = json.Unmarshal(c.GetRequestBody(), &form); err == nil {
		validation := validation.Validation{}
		if valid, _ := validation.Valid(&form); !valid {
			c.Error(err)
		} else if data, err = c.Component.TransferAmount(form); err == nil {
			c.Data["json"] = requestresponseutils.PrepareResponse(data, err, "Success")
			c.ServeJSON()
		}
	}
	c.Error(err)
}
