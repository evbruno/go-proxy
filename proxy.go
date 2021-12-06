// eg: go run proxy.go -url http://localhost:8443
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

var requestTimeout uint = 0

type Prox struct {
	target *url.URL
	proxy  *httputil.ReverseProxy
}

func NewProxy(target string) *Prox {
	url, _ := url.Parse(target)
	return &Prox{target: url, proxy: httputil.NewSingleHostReverseProxy(url)}
}

func (p *Prox) handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-GoProxy", "GoProxy")
	p.proxy.Transport = &myTransport{}
	p.proxy.ServeHTTP(w, r)
}

type myTransport struct {
}

func (t *myTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	fmt.Println("Request received... sleeping for (sec)", requestTimeout)
	time.Sleep(time.Duration(requestTimeout) * time.Second)

	body, err := httputil.DumpRequestOut(request, true)

	if err != nil {
		print("What?", err)
		return nil, err
	}

	fmt.Println("Request Body : ", string(body))

	response, err := http.DefaultTransport.RoundTrip(request)
	if err != nil {
		print("\n\ncame in error resp here", err)
		return nil, err //Server is not reachable. Server not working
	}

	// body, err := httputil.DumpResponse(response, true)
	// if err != nil {
	//     print("\n\nerror in dumb response")
	//     // copying the response body did not work
	//     return nil, err
	// }

	log.Println("Response Body : ", string(body))
	return response, err
}

func main() {
	const (
		defaultPort         = ":9090"
		defaultPortUsage    = "default server port, ':9090'"
		defaultTarget       = "http://127.0.0.1:8080"
		defaultTargetUsage  = "default redirect url, 'http://127.0.0.1:8080'"
		defaultTimeout      = 0
		defaultTimeoutUsage = "default timeout in seconds, '0'"
	)

	// flags
	port := flag.String("port", defaultPort, defaultPortUsage)
	redirecturl := flag.String("url", defaultTarget, defaultTargetUsage)

	flag.UintVar(&requestTimeout, "timeout", defaultTimeout, defaultTimeoutUsage)

	flag.Parse()

	fmt.Println("server will run on :", *port)
	fmt.Println("redirecting to :", *redirecturl)

	if requestTimeout > 0 {
		fmt.Println("timeout set to :", requestTimeout)
	}

	// proxy
	proxy := NewProxy(*redirecturl)

	//http.HandleFunc("/proxyServer", ProxyServer)

	// server redirection
	http.HandleFunc("/", proxy.handle)
	log.Fatal(http.ListenAndServe(*port, nil))
}
