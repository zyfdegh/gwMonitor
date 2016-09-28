package main

import (
	"log"
	"monitor/services"
	"net/http"

	"github.com/emicklei/go-restful"
)

func main() {

	wsContainer := restful.NewContainer()
	u := services.Resource{}
	u.Register(wsContainer)

	log.Printf("start listening on localhost:18089")
	server := &http.Server{Addr: ":18089", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
