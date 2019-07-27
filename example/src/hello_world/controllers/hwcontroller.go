package controllers

import "github.com/astaxie/beego"

type HWController struct {
	beego.Controller
}

type JSONStruct struct {
	Code int
	Msg  string
	Data string
}

//func (hw *HWController) Get() {      // get请求
//	hw.TplName = "hello world.html"  // 关联静态文件
//}

func (hw *HWController) Get() {      // get请求
	json := &JSONStruct{0, "Success", "hello json"}
	hw.Data["json"] = json
	hw.ServeJSON()
	return
}
