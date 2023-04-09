package main

import (
	"fmt"
	"haste/infra/driver/routers"
	"log"
	"os"

	_ "haste/adapters/repositories/account"
	_ "haste/adapters/repositories/user"
	_ "haste/core/components/account"
	_ "haste/core/components/user"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	// Generated using http://patorjk.com/software/taag/#p=display&f=Graffiti
	log.Println(`
    ___  ___ ___  _ __ __  _ __ __   ___  _  __ ___ __            __ _ _ __ (_)
   / _ \/ __/ _ \| '_  _ \| '_  _ \ / _ \| '__/ __/ _ \  _____   / _  | |\ \| |
  |  __/ (_| (_) | | | | || | | | ||  __/| | | (_|  __/ |_____| | (_| | |_) | |
   \___|\___\___/|_| |_| ||_| |_| | \___||_|  \___\___|          \__,_|  __/|_|
                                                                      |_|    
    `)

	// Initialize app
	AppInit()

	// GET http://localhost:8080/health => ctrl.Health()

	// web.Router("/health", &user.UserController{}, "get:Health")
	web.Run()
}

func AppInit() {
	routers.InitRoutes()

	// if beego.BConfig.RunMode == "dev" {
	// 	beego.BConfig.WebConfig.DirectoryIndex = true
	// 	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	// }
	appName, _ := web.AppConfig.String("appname")
	sqlconn, _ := web.AppConfig.String("sqlconn")
	db := os.Getenv("DB_HOST")
	fmt.Println(appName, "appname", db, sqlconn)

	// serviceName, _ = web.AppConfig.String("ServiceName")

	// // Initializing global db
	// dbConfig, _ := web.AppConfig.DIY("db_config")
	// var dbConfigMap map[string]string
	// _ = maputils.MapInterfaceToObject(dbConfig, &dbConfigMap)

	// if err := global.InitAllDB(dbConfigMap); err != nil {
	// 	logger.LogErr(gocontext.Background(), "error", err)
	// }
}
