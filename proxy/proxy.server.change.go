package proxy

import (
	"fmt"

	"github.com/micro-plat/hydra/conf"
	"github.com/micro-plat/hydra/servers/pkg/circuit"
)

//SetTrace 显示跟踪信息
func (s *ProxyServer) SetTrace(b bool) {
	s.conf.SetMetadata("show-trace", b)
	return
}

//SetMetric 重置metric
func (s *ProxyServer) SetMetric(metric *conf.Metric) error {
	s.metric.Stop()
	if metric.Disable {
		return nil
	}
	if err := s.metric.Restart(metric.Host, metric.DataBase, metric.UserName, metric.Password, metric.Cron, s.Logger); err != nil {
		err = fmt.Errorf("metric设置有误:%v", err)
		return err
	}
	return nil
}

//SetHeader 设置http头
func (s *ProxyServer) SetHeader(headers conf.Headers) error {
	s.conf.SetMetadata("headers", headers)
	return nil
}

//StopMetric stop metric
func (s *ProxyServer) StopMetric() error {
	s.metric.Stop()
	return nil
}

//CloseCircuitBreaker 关闭熔断配置
func (s *ProxyServer) CloseCircuitBreaker() error {
	if c, ok := s.conf.GetMetadata("__circuit-breaker_").(*circuit.NamedCircuitBreakers); ok {
		c.Close()
	}
	return nil
}

//SetCircuitBreaker 设置熔断配置
func (s *ProxyServer) SetCircuitBreaker(c *conf.CircuitBreaker) error {
	s.conf.SetMetadata("__circuit-breaker_", circuit.NewNamedCircuitBreakers(c))
	return nil
}
