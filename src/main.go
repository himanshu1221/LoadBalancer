package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

// step 4 Creating server server interface
type Server interface {
	Address() string
	IsAlive() bool
	Server(rw http.ResponseWriter, r *http.Request)
}

// Step 1 created struct
type simpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}

// Created a function for handdling the error
func handleErr(err error) {
	if err != nil {
		fmt.Println("ERR")
		os.Exit(1)
	}
}

// Step 3 creating a struct for loadbalancer
type LoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []Server
}

// Step 2 intialising the server
func newSimpleServer(addr string) *simpleServer {
	serverUrl, err := url.Parse(addr)
	handleErr(err)
	return &simpleServer{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}
