package main

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/hydra"
	"github.com/micro-plat/hyproxy/proxy/middleware"
	"github.com/urfave/cli"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func (app *hyproxy) init() {

	app.Cli.Append(hydra.ModeRun, cli.StringSliceFlag{
		Name:  "filter,f",
		Usage: "设置域名过滤",
	})

	app.Initializing(func(c component.IContainer) error {
		middleware.Hosts = app.Cli.Context().StringSlice("filter")
		return nil
	})
}
