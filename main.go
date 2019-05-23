package main

import (
	"coterey/routers"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

func main() {
	routers.InitApp()
	beego.Router("/", &routers.HomeRouter{})
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.SetStaticPath("/.well-known/brave-rewards-verification.txt", "static/brave-rewards-verification.txt")
	beego.SetStaticPath("/cv-en.pdf", "static/cv-en.pdf")
	beego.SetStaticPath("/cv-es.pdf", "static/cv-es.pdf")
	beego.Run()
}
