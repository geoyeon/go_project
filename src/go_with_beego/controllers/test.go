package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

type jsonstruct struct {
  FieldOne string `json:"field_one"`
}

func (c *TestController) Get() {
	data := map[string]interface{}{
  "test":  "test",
  "email":   "geoyeon.gmail.com",
  "array": []string{"one", "two"},
 }
 c.Data["json"] = &data
 c.ServeJSON()
}
