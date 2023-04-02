package controllers

import (
	"context"
	"haste/core/components"
	"haste/pkg/types"
	"haste/pkg/utils"
	requestresponseutils "haste/pkg/utils/request_response"

	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	web.Controller
	AccessToken string
	ReqCtx      context.Context
	Errors      *types.Errors
}

// Preparer
type Preparer interface {
	UpdateComponent(interface{})
}

// Prepare is called before the http action is processes, to initialize
func (bc *BaseController) Prepare() {
	fmt.Println("prepare ")
	requestCtx := bc.Ctx.Request.Context()
	bc.ReqCtx = requestCtx
	bc.Errors = new(types.Errors)

	// srv, _ := bc.ReqCtx.Value(types.Key("srv")).(string)

	// controller, action := bc.GetControllerAndAction()
	// spanName := fmt.Sprintf("%v.%v", controller, action)
	// eapm, _ := bc.ReqCtx.Value(types.Key("eapm")).(*elastic.ElasticAPM)
	// eapm.StartSpan(spanName, "controller", bc.ReqCtx)

	bc.AccessToken, _ = bc.ReqCtx.Value(types.Key("access_token")).(string)

	// var sentryScope = make(map[string]interface{})

	// bc.ReqCtx = context.WithValue(bc.ReqCtx, types.Key("sentry_scope"), sentryScope)

	if app, ok := bc.AppController.(Preparer); !ok {
		fmt.Println("do nothing")
		// do nothing
	} else if component, err := bc.InitComponent(); err != nil {
		bc.Error(err)
	} else {
		app.UpdateComponent(component)
	}
}

// Finish is executed after the related http method, to clean up
func (bc *BaseController) Finish() {
	// eapm, _ := bc.ReqCtx.Value(types.Key("eapm")).(*elastic.ElasticAPM)
	// eapm.EndSpan()
	fmt.Println("finish in base controller")
}

// InitComponent initializes the component whose method needs to be called.
// It returns error.
func (bc *BaseController) InitComponent() (interface{}, error) {
	controller, _ := bc.GetControllerAndAction()
	fmt.Println("base controller - current controller", controller)
	componentKey := strings.Replace(controller, "Controller", "", -1)
	fmt.Println("compo key", componentKey)

	base := &components.BaseComponent{
		AccessToken: bc.AccessToken,
		ReqCtx:      bc.ReqCtx,
		Headers:     make(utils.HType),
		Errors:      bc.Errors,
		Extra: map[string]interface{}{
			"ReqIP":        bc.Ctx.Input.IP(),
			"ReqTimestamp": time.Now().Format(time.RFC3339),
		},
	}
	for k := range bc.Ctx.Request.Header {
		base.Headers[strings.ToLower(k)] = bc.Ctx.Request.Header.Get(k)
	}
	componentFn, ok := components.ComponentMap[componentKey]
	if !ok {
		err := fmt.Errorf("failed to initialize component: %v", componentKey)
		return nil, err
	}

	return componentFn(base), nil
}

// Error is used to stop execution, if any fatal error has occured.
func (bc *BaseController) Error(err error) {

	fmt.Println("Error cam", err)
	bc.Data["json"] = requestresponseutils.PrepareResponse(nil, err, "Something went wrong!")
	bc.Ctx.Output.SetStatus(http.StatusInternalServerError)
	bc.ServeJSON()
	bc.Finish()  // need to call finish method manually before stopping execution
	bc.StopRun() // stop controller execution immediately
}

// AddHeaders adds additional response headers.
func (bc *BaseController) AddHeaders(status int, opts map[string]bool) {
	if fl, ok := opts["no_cache"]; ok && fl {
		bc.Ctx.Output.Header("Cache-Control", "no-store, max-age=0")
	}

	if os.Getenv("ENABLE_HSTS") == "true" {
		// set max-age of 1 year
		bc.Ctx.Output.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
	}

	bc.Ctx.Output.SetStatus(status)
}

// this function is used to set http cookie in response.
func (bc *BaseController) SetCookies(cookies *[]http.Cookie) {
	for _, c := range *cookies {
		if os.Getenv("SECURE_COOKIE") == "true" {
			c.Secure = true
		}
		http.SetCookie(bc.Ctx.ResponseWriter, &c)
	}
}

// GetRequestBody fetches the body of the incoming request.
// It returns the body byte data.
func (bc *BaseController) GetRequestBody() []byte {
	bodyBytes := bc.Ctx.Input.RequestBody
	if decodedBytes, err := base64.StdEncoding.DecodeString(string(bodyBytes)); err == nil {
		bodyBytes = decodedBytes
	}
	if string(bodyBytes) == "" {
		bodyBytes = []byte(`{}`)
	}
	return bodyBytes
}

func init() {
	fmt.Println("init base controller")
}
