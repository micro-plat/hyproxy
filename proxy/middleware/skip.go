package middleware

import (
	"fmt"
	"strings"

	"github.com/elazarl/goproxy"
	"github.com/micro-plat/hydra/conf"
)

var Hosts = []string{}

func skip(conf *conf.MetadataConf, ctx *goproxy.ProxyCtx, header map[string][]string) bool {
	if getSkip(ctx) {
		return true
	}
	//服务名过滤
	chost := fmt.Sprintf("%s://%s%s", ctx.Req.URL.Scheme, ctx.Req.URL.Host, ctx.Req.URL.Path)
	for _, h := range Hosts {
		if strings.Contains(chost, h) {
			return false
		}
	}
	return true

}
