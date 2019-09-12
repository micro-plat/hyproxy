package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/elazarl/goproxy"
	"github.com/micro-plat/hydra/conf"
	"github.com/micro-plat/lib4go/logger"
)

type RequestFunc func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response)
type ResponseFunc func(r *http.Response, ctx *goproxy.ProxyCtx) *http.Response

func setLogger(c *goproxy.ProxyCtx, l *logger.Logger) {
	c.UserData = make(map[string]interface{})
	c.UserData.(map[string]interface{})["__logger_"] = l
}
func getLogger(c *goproxy.ProxyCtx) *logger.Logger {
	data := c.UserData.(map[string]interface{})
	return data["__logger_"].(*logger.Logger)
}
func setStartTime(c *goproxy.ProxyCtx) {
	c.UserData.(map[string]interface{})["__start_"] = time.Now()
}
func getStartTime(c *goproxy.ProxyCtx) time.Time {
	return c.UserData.(map[string]interface{})["__start_"].(time.Time)
}
func setSkip(c *goproxy.ProxyCtx) {
	c.UserData.(map[string]interface{})["__skip__"] = true
}
func getSkip(c *goproxy.ProxyCtx) bool {
	if v, ok := c.UserData.(map[string]interface{})["__skip__"]; ok {
		return v.(bool)
	}
	return false
}

//LoggingHead 记录日志
func LoggingHead(conf *conf.MetadataConf) RequestFunc {
	return func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		log := logger.New("proxy")
		setLogger(ctx, log)
		setStartTime(ctx)
		path := fmt.Sprintf("%s://%s%s", r.URL.Scheme, r.URL.Host, r.URL.Path)
		log.Info("proxy.request", r.Method, path, "from", r.RemoteAddr)
		return r, nil
	}
}

func LoggingTail(conf *conf.MetadataConf) ResponseFunc {
	return func(r *http.Response, ctx *goproxy.ProxyCtx) *http.Response {

		start := getStartTime(ctx)
		statusCode := 0
		msg := ""
		high := ""
		if r != nil {
			statusCode = r.StatusCode
		} else if ctx.Error != nil {
			msg = ctx.Error.Error()
		}
		if time.Since(start) > time.Second {
			high = "⇡"
		}
		path := fmt.Sprintf("%s://%s%s", ctx.Req.URL.Scheme, ctx.Req.URL.Host, ctx.Req.URL.Path)
		if statusCode >= 200 && statusCode < 400 {
			getLogger(ctx).Info("proxy.response", ctx.Req.Method, path, statusCode, time.Since(start), high)
		} else {
			getLogger(ctx).Error("proxy.response", ctx.Req.Method, path, statusCode, msg, time.Since(start), high)
		}
		getLogger(ctx).Close()
		return r
	}
}
