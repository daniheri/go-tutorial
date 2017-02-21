package controllers

import (
	"github.com/astaxie/beego"
	"daniapi/models"
	"encoding/json"
	"net/http"
)

type LatihanController struct {
	beego.Controller
}

// @Title dani
// @Description ini dani latihan
// @Success 200 {string} latihan success
// @router / [get]
func (u *LatihanController) Dani() {

	fromModel := models.Latihan{Nama: "Dani heriyanto",
		Kelas: "12 TKJ", Gender: getGender(1),
		Tanggal: "12-12-1992",
		Hobi: [] models.Category {{"mancing", "madang"}},
	}

	u.Data["json"] = &fromModel
	u.ServeJSON()
}

// @Title dani
// @Description ini dani latihan
// @Success 200 {string} latihan success POST
// @router / [post]
func (data *LatihanController) DaniPost() {

	var requestBody models.LatihanRequest
	var output models.JsonErrorr
	var header = data.Ctx.Input.Header("ID")

	json.Unmarshal(data.Ctx.Input.RequestBody, &requestBody)

	fromModel := models.Latihan{
		Nama: "Dani heriyanto POST",
		Kelas: "12 TKJ POST",
		Gender: getGender(1),
		Tanggal: "12-12-1992 POST",
		Hobi: []models.Category{{"mbuh", "yuhu"}},
	}

	if requestBody.Token == "" || header == ""{

		output.Output = "you errorr"
		data.Data["json"] = &output
		data.Ctx.Output.SetStatus(http.StatusBadRequest)

	} else {
		data.Data["json"] = &fromModel
	}

	data.ServeJSON()
}

func getGender(gender int) string {
	if gender == 1{
		return "laki-laki"
	} else {
		return "terserah"
	}
}
