package repositories

import (
	"context"
	"fmt"
)

type BaseRepository struct {
	ReqCtx context.Context
}

var RepositoryMap = make(map[string]func(*BaseRepository) interface{})

func (be *BaseRepository) Prepare() {
	fmt.Println("prepare base entity go")
}

func init() {
	fmt.Println("init of base entities")
}
