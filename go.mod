module github.com/micro-plat/hyproxy

go 1.12

require (
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a
	github.com/elazarl/goproxy v0.0.0-20190711103511-473e67f1d7d2
	github.com/elazarl/goproxy/ext v0.0.0-20190910210725-4d0f5f06fe9d // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/micro-plat/hydra v0.10.15
	github.com/micro-plat/lib4go v0.1.8
	github.com/urfave/cli v1.20.0
	github.com/zkfy/log v0.0.0-20180312054228-b2704c3ef896
)

replace github.com/micro-plat/hydra => ../../../github.com/micro-plat/hydra
