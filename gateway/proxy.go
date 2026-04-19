package gateway

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	if !Authorized(r) {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	ip := r.RemoteAddr

	if !Allow(ip) {
		http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
		return
	}

	Log(r)

	targetURL := getBackend(r.URL.Path)
	target, _ := url.Parse(targetURL)

	proxy := httputil.NewSingleHostReverseProxy(target)

	originalPath := r.URL.Path

	for prefix := range routeMap {
		if strings.HasPrefix(originalPath, prefix) {
			r.URL.Path = strings.TrimPrefix(originalPath, prefix)
			if r.URL.Path == "" {
				r.URL.Path = "/"
			}
			break
		}
	}

	r.Host = target.Host

	proxy.ServeHTTP(w, r)
}
