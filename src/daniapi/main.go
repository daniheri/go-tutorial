package main

import(
	_ "daniapi/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

func main() {
	// Create the database handle, confirm driver is present
	db, _ := sql.Open("mysql", "dellis:@/shud")
	defer db.Close()
	// Connect and check the server version
	 var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)


	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
