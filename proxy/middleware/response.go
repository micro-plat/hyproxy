package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/elazarl/goproxy"
	"github.com/micro-plat/hydra/conf"
	"github.com/micro-plat/lib4go/types"
)

func Response(conf *conf.MetadataConf) ResponseFunc {
	return func(r *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
		if r == nil || skip(conf, ctx, r.Header) {
			return r
		}
		header := map[string][]string{}
		if r != nil {
			header = r.Header
		}
		headerBuff := make([]string, 0, len(header))
		for h, v := range header {
			headerBuff = append(headerBuff, fmt.Sprintf("%s:\t%v", h, strings.Join(v, ",")))
		}

		bodyBytes, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		logger := getLogger(ctx)

		logger.Infof(
			"\n"+
				`------------------------response-----------------------------------------------------
+General
	Request URL:	%s
	Request Method:	%s
	Status Code:	%d
	Remote Address:	%s
+Response Headers
	%s
+Raw
	%s
---------------------------------------------------------------------------------------`, ctx.Req.URL,
			ctx.Req.Method,
			r.StatusCode,
			ctx.Req.RemoteAddr,
			strings.Join(headerBuff, "\n\t"),
			types.GetString(string(bodyBytes), "-"))

		return r
	}
}
