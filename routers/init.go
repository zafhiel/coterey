package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/beego/i18n"
	"strings"
)

func InitApp() {
	initLocales()
}

func initLocales() {
	// Initialized language type list.
	langs := strings.Split(beego.AppConfig.String("lang::types"), "|")
	names := strings.Split(beego.AppConfig.String("lang::names"), "|")
	langTypes = make([]*langType, 0, len(langs))
	for i, v := range langs {
		langTypes = append(langTypes, &langType{
			Lang: v,
			Name: names[i],
		})
	}

	for _, lang := range langs {
		logs.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			logs.Error("Fail to set message file: " + err.Error())
			return
		}
	}
}
