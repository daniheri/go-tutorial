package controllers

import (
	"github.com/astaxie/beego"
	"gototo/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) GetPeople() {
	people, err := models.GetPeople()
	if err != nil {
		c.Data["json"] = err
		c.ServeJSON()
		return
	}

	c.Data["json"] = people
	c.ServeJSON()
	return
}
