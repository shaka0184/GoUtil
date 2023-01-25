package httpUtil

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type ProxyTarget struct {
	Prefix string
	URL    string
}

func PathMatchReverseProxy(target []ProxyTarget) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		requestPath := r.URL.Path

		for _, v := range target {
			pre := "/" + v.Prefix + "/"
			if strings.HasPrefix(requestPath, pre) {
				targetUrl, err := url.Parse(v.URL)
				if err != nil {
					log.Fatal(err)
				}

				r.URL.Path = strings.TrimPrefix(requestPath, pre)
				proxy := httputil.NewSingleHostReverseProxy(targetUrl)
				proxy.ServeHTTP(w, r)
			}
		}
	}
}
