package services

import (
	"encoding/json"
	"errors"
	"log"
	"net"
	"os"
	"strings"

	"github.com/jmoiron/jsonq"
)

func UdpCall(server, msg string) (info string, err error) {

	addr, err := net.ResolveUDPAddr("udp", server)
	if err != nil {
		log.Println("Can't resolve address: ", err)
		return "", err
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Println("Can't dial: ", err)
		return "", err
	}

	defer conn.Close()

	log.Println("Writing something to server...")
	_, err = conn.Write([]byte(msg))
	if err != nil {
		log.Println("failed:", err)
		return "", err
	}

	data := make([]byte, 1024)

	n, remoteAddr, err := conn.ReadFromUDP(data)
	log.Println("Connecting...")
	if err != nil {
		log.Println("failed to read UDP msg because of ", err)
		return "", err
	}

	if remoteAddr != nil {
		log.Println("got message from ", remoteAddr, " with n = ", n)
		if n > 0 {
			log.Println("from address", remoteAddr, "got message:", string(data[0:n]), n)
			info = string(data[0:n])
		}
	}

	log.Println(info)
	return info, nil
}

func getAddrs() (addrs []string, err error) {

	strAddrs := os.Getenv("ADDRESSES")
	if strings.EqualFold(strAddrs, "nil") {
		err = errors.New("getAddrs failed, find no addresses")
		return
	}
	//strAddrs looks like : "127.0.0.1:8080,127.0.0.1:8081"
	addrs = strings.Split(strAddrs, ",")
	return
}

func getMonitorType() (mtype string) {

	//PGW or SGW
	mtype = os.Getenv("MONITOR_TYPE")
	return
}

func parseJson(jsonstring string) (instances, connNum, ovsId int) {

	data := map[string]interface{}{}
	result := json.NewDecoder(strings.NewReader(jsonstring))
	result.Decode(&data)
	jq := jsonq.NewQuery(data)
	instances, _ = jq.Int("instances")
	connNum, _ = jq.Int("connNum")
	ovsId, _ = jq.Int("ovsId")
	return
}

func process(infos []string) (sumInstance int, sumConn int, monitorType string, err error) {

	sumInstance = 0
	sumConn = 0
	//get sumInstance and sumConn
	for _, info := range infos {
		instances, connNum, _ := parseJson(info)
		sumInstance = sumInstance + instances
		sumConn = sumConn + connNum
	}
	monitorType = getMonitorType()
	log.Println("sumInstance=", sumInstance, ", sumConn=", sumConn, ", monitorType=", monitorType)
	return
}
