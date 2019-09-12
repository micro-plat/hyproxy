package proxy

import (
	"github.com/micro-plat/hydra/conf"
	"github.com/micro-plat/hydra/servers"
	"github.com/micro-plat/lib4go/logger"
)

type proxyServerAdapter struct {
}

func (h *proxyServerAdapter) Resolve(registryAddr string, conf conf.IServerConf, log *logger.Logger) (servers.IRegistryServer, error) {
	return NewApiResponsiveServer(registryAddr, conf, log)
}

func init() {
	servers.Register("proxy", &proxyServerAdapter{})
}
