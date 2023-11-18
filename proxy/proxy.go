package proxy

import (
	"net/http"
	"net/http/httputil"
	"time"

	"k8s.io/client-go/kubernetes"
)

type Proxy struct {
	k8sClient kubernetes.Interface
	rp        *httputil.ReverseProxy
}

func NewProxy(k8sClient kubernetes.Interface) *Proxy {
	rp := &httputil.ReverseProxy{
		ModifyResponse: func(*http.Response) error {
			// todo: need to do anything here?
			return nil
		},
		Transport: &http.Transport{
			TLSHandshakeTimeout:   10 * time.Second,
			DisableKeepAlives:     false,
			DisableCompression:    true,
			IdleConnTimeout:       1 * time.Minute,
			ResponseHeaderTimeout: 30 * time.Second,
		},
	}

	p := &Proxy{
		k8sClient: k8sClient,
		rp:        rp,
	}
	rp.Rewrite = p.rewrite
	return p
}

func (p *Proxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// todo: check auth here
	//action, permission, err := currentAction(req)

	p.rp.ServeHTTP(rw, req)
}

func (p *Proxy) rewrite(*httputil.ProxyRequest) {
	// todo: attach downstream auth and rewrite request
}
