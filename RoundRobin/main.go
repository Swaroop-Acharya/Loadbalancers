package main

import (
	"internal/runtime/strconv"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

type Server interface {
	GetAddr() string
	IsAlive() bool
	Serve(w http.ResponseWriter, r *http.Request)
}

type LoadBalancer struct {
	Port           string 
	Servers         []Server
	RoundRobinCount int
}

type SimpleServer struct {
	Addr  string
	Proxy *httputil.ReverseProxy
}

func HandleError(err error) {
	log.Fatalf("Error occured: %v \n", err)
	os.Exit(1)
}

func newLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		Port:            port,
		Servers:         servers,
		RoundRobinCount: 0,
	}
}

func newSimpleServer(addr string) *SimpleServer {
	serverURL, err := url.Parse(addr)
	HandleError(err)
	return &SimpleServer{
		Addr:  addr,
		Proxy: httputil.NewSingleHostReverseProxy(serverURL),
	}
}

func (ns *SimpleServer) GetAddr() string { return ns.Addr }
func (ns *SimpleServer) IsAlive() bool   { return true }
func (ns *SimpleServer) Server(w http.ResponseWriter, r *http.Request) {
	ns.Proxy.ServeHTTP(w, r)
}

func main() {
	servers := []Server{
		newSimpleServer("https://www.facebook.com"),
		newSimpleServer("https://www.bing.com"),
		newSimpleServer("https://www.duckduckgo.com"),
	}
	lb := newLoadBalancer(8000, servers)
	
	http.ListenAndServe( , handleRedirect)

}
