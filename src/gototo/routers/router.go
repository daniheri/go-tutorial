package routers

import (
	"gototo/controllers"
	"github.com/astaxie/beego"
)

func init() {

	news := beego.NewNamespace("/v1",
		beego.NSRouter("/people", &controllers.MainController{}, "get:GetPeople"),
	)

	beego.AddNamespace(news)
}
