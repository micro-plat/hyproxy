package http

import (
	"context"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"github.com/micro-plat/hydra/component"
)

//ProxyHandler 正向代理
type ProxyHandler struct {
	container component.IContainer
}

//NewProxyHandler 构建正向代理
func NewProxyHandler(container component.IContainer) (u *ProxyHandler) {
	return &ProxyHandler{container: container}
}

// func (u *ProxyHandler) Handle(ctx *context.Context) (r interface{}) {
// 	remote, err := url.Parse("http://127.0.0.1:443")
// 	if err != nil {
// 		panic(err)
// 	}
// 	proxy := httputil.NewSingleHostReverseProxy(remote)
// 	var pTransport http.RoundTripper = &http.Transport{
// 		Proxy:                 http.ProxyFromEnvironment,
// 		Dial:                  Dial,
// 		TLSHandshakeTimeout:   10 * time.Second,
// 		ExpectContinueTimeout: 1 * time.Second,
// 	}
// 	proxy.Transport = pTransport

// 	proxy.ServeHTTP(w, r)
// }

//Handle 正向代理处理函数
func (u *ProxyHandler) XHandle(ctx *context.Context) (r interface{}) {

	// step 1
	req, err := ctx.Request.Http.Get()
	if err != nil {
		return err
	}
	if strings.ToUpper(ctx.Request.GetMethod()) == "CONNECT" {
		return
	}

	outReq := new(http.Request)
	transport := http.DefaultTransport
	*outReq = *req // this only does shallow copies of maps

	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		if prior, ok := outReq.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ", ") + ", " + clientIP
		}
		outReq.Header.Set("X-Forwarded-For", clientIP)
	}

	// step 2
	res, err := transport.RoundTrip(outReq)
	if err != nil {
		ctx.Response.SetStatus(http.StatusBadGateway)
		return
	}

	// step 3
	for key, value := range res.Header {
		for _, v := range value {
			ctx.Response.SetHeader(key, v)
		}
	}

	ctx.Response.SetStatus(res.StatusCode)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return body

}
