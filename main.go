package main

import (
	"log"
	"net/http"

	"github.com/LinkerNetworks/gwMonitor/autoscaling"
	"github.com/LinkerNetworks/gwMonitor/services"
	"github.com/emicklei/go-restful"
)

func main() {
	// startRestfulServer()
	autoscaling.Start()
}

func startRestfulServer() {
	wsContainer := restful.NewContainer()
	u := services.Resource{}
	u.Register(wsContainer)

	log.Printf("start listening on localhost:18089")
	server := &http.Server{Addr: ":18089", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
