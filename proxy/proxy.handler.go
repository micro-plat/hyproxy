package proxy

import (
	"net/http"

	"github.com/elazarl/goproxy"
	"github.com/micro-plat/hydra/conf"
	"github.com/micro-plat/hyproxy/proxy/middleware"
)

type handler struct {
	engine *goproxy.ProxyHttpServer
}

func newHandler(conf *conf.MetadataConf) *handler {
	engine := goproxy.NewProxyHttpServer()
	engine.OnRequest().DoFunc(middleware.LoggingHead(conf))
	engine.OnRequest().DoFunc(middleware.Request(conf))

	engine.OnResponse().DoFunc(middleware.Response(conf))
	engine.OnResponse().DoFunc(middleware.LoggingTail(conf))
	return &handler{
		engine: engine,
	}
}
func (h *handler) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, h.engine)
}
