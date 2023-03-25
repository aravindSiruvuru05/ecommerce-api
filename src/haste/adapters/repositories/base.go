package repositories

import (
	"context"
	"fmt"
)

type BaseEntity struct {
	ReqCtx context.Context
}

var EntityMap = make(map[string]func(*BaseEntity) interface{})

func (be *BaseEntity) Prepare() {
	fmt.Println("prepare base entity go")
}

func init() {
	fmt.Println("init of base entities")
}
