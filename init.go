package main

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hyproxy/services/http"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func (app *hyproxy) init() {

	app.install()

	app.handling()

	app.Initializing(func(c component.IContainer) error {
		return nil
	})
	app.API("/*name", http.NewProxyHandler)
}
