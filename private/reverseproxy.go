package private

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ServeReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	tUrl, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(tUrl)

	req.URL.Host = tUrl.Host
	req.URL.Scheme = tUrl.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = tUrl.Host

	proxy.ServeHTTP(res, req)
}
