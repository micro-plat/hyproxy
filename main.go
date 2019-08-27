package main

import "github.com/micro-plat/hydra/hydra"
import _ "github.com/go-sql-driver/mysql"

type hyproxy struct {
	*hydra.MicroApp
}

func main() {
	app := &hyproxy{
		hydra.NewApp(
			hydra.WithPlatName("hyproxy"),
			hydra.WithSystemName("proxy"),
			hydra.WithServerTypes("api"),
			hydra.WithDebug()),
	}

	app.init()

	app.Start()
}
