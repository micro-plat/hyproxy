package main

import "github.com/micro-plat/hydra/hydra"
import _ "github.com/go-sql-driver/mysql"
import _ "github.com/micro-plat/hyproxy/proxy"

type hyproxy struct {
	*hydra.MicroApp
}

func main() {
	app := &hyproxy{
		hydra.NewApp(
			hydra.WithPlatName("hyproxy"),
			hydra.WithSystemName("hyproxy"),
			hydra.WithServerTypes("proxy"),
			hydra.WithClusterName("proxy")),
	}

	app.init()

	app.Start()
}
