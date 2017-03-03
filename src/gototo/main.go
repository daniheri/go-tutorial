package main
import (
	_ "gototo/routers"
	"github.com/astaxie/beego"
	"gototo/db"
)

func main() {
	beego.Run()
	db.Conn()
}