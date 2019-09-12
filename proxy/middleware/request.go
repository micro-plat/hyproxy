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

func Request(conf *conf.MetadataConf) RequestFunc {
	return func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		if skip(conf, ctx, r.Header) {
			return r, nil
		}
		logger := getLogger(ctx)
		headerBuff := make([]string, 0, len(r.Header))
		for h, v := range r.Header {
			headerBuff = append(headerBuff, fmt.Sprintf("%s:\t%v", h, strings.Join(v, ",")))
		}
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		logger.Infof(
			"\n"+
				`------------------------request-----------------------------------------------------
+General
	Request URL:	%s
	Request Method:	%s
	Remote Address:	%s
+Request Headers
	%s
+Raw
	%s
-------------------------------------------------------------------------------------`, r.URL,
			r.Method,
			r.RemoteAddr,
			strings.Join(headerBuff, "\n\t"),
			types.GetString(string(bodyBytes), "-"))

		return r, nil
	}
}
