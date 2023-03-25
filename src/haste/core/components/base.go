package components

import (
	"context"
	"fmt"
	"haste/adapters/repositories"
	"haste/pkg/types"
	"haste/pkg/utils"
	"net/http"
)

type BaseComponent struct {
	AccessToken string                 `json:"-"`
	ReqCtx      context.Context        `json:"-"`
	Headers     utils.HType            `json:"-"`
	Cookies     map[string]http.Cookie `json:"-"`
	Errors      *types.Errors          `json:"-"`
	Extra       map[string]interface{} `json:"-"`
}

// ComponentMap
var ComponentMap = make(map[string]func(*BaseComponent) interface{})

func (bc *BaseComponent) GetEntityObject(entityName string) interface{} {
	return repositories.EntityMap[entityName](&repositories.BaseEntity{ReqCtx: bc.ReqCtx})
}

func (bc *BaseComponent) CreateBaseEntity() repositories.BaseEntity {
	return repositories.BaseEntity{ReqCtx: bc.ReqCtx}
}

func init() {
	fmt.Println("init base component")
}
