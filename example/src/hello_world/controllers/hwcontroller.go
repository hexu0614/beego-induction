package controllers

import "github.com/astaxie/beego"

type HWController struct {
	beego.Controller
}

func (hw *HWController) Get() {      // get请求
	hw.TplName = "hello world.html"  // 关联静态文件
}

