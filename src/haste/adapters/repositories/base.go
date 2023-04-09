package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	db "haste/infra/driven/database/sqlc"

	_ "github.com/lib/pq"
)

type BaseRepository struct {
	ReqCtx context.Context
}

var DBQueries *db.Queries

var RepositoryMap = make(map[string]func(*BaseRepository) interface{})

func (be *BaseRepository) Prepare() {
	fmt.Println("prepare base entity go")
}

func init() {
	DBDriver := "postgres"
	DBSource := "postgresql://user:user@db:5432/haste?sslmode=disable"

	conn, err := sql.Open(DBDriver, DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	} else {
		log.Println("Successfully connected to db.")
	}

	DBQueries = db.New(conn)

	fmt.Println("init of base entities")
}
