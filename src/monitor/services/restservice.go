package services

import (
	"log"
	"strings"

	"github.com/emicklei/go-restful"
)

type Resource struct {
}

type RespStruct struct {
	Success bool
	Data    interface{}
	Err     string
}

type RespData struct {
	Instances   int
	ConnNum     int
	MonitorType string
}

func (r Resource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/monitor").
		Doc("monitor").
		Consumes("*/*").
		Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(ws.GET("/").To(r.callServers).
		// docs
		Doc("monitor"))

	container.Add(ws)
}

func (r Resource) callServers(request *restful.Request, response *restful.Response) {

	log.Println("callservers...")

	//get UDP server addresses from ENV file
	addrs, err := getAddrs()
	if err != nil {
		resp := RespStruct{Success: false, Err: err.Error()}
		response.WriteAsJson(resp)
		return
	}
	log.Println(addrs, len(addrs))

	monitorType, err := getMonitorType()
	if err != nil {
		log.Println(monitorType)
		resp := RespStruct{Success: false, Err: err.Error()}
		response.WriteAsJson(resp)
		return
	}
	log.Println("MONITOR_TYPE: ", monitorType)

	infos := make([]string, 0, len(addrs))

	//call UDP servers
	for _, address := range addrs {
		info, err := UdpCall(strings.TrimSpace(address), "hi")
		if err != nil {
			log.Println("UdpCall "+strings.TrimSpace(address)+" failed.", err)
		}
		info = strings.TrimSpace(info)
		infos = append(infos, info)
	}

	log.Println(infos)

	instances, connNum, _ := process(infos)
	respData := RespData{Instances: instances, ConnNum: connNum, MonitorType: monitorType}
	resp := RespStruct{Success: true, Data: respData}
	response.WriteAsJson(resp)
}
