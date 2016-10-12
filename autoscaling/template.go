package autoscaling

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/LinkerNetworks/gwMonitor/conf"
	marathon "github.com/gambol99/go-marathon"
)

var (
	onceTemplate sync.Once

	jsonPath   string
	pgwGroupID string
	sgwGroupID string
)

func init() {
	onceTemplate.Do(func() {
		monitorType := env(keyMonitorType).Value
		switch monitorType {
		case typePGW:
			jsonPath = conf.OptionsReady.PgwJSON
			pgwGroupID = groupID()
		case typeSGW:
			jsonPath = conf.OptionsReady.SgwJSON
			sgwGroupID = groupID()
		default:
			log.Printf("unknow monitor type: %s\n", monitorType)
		}
	})
}

// return apps from 0 to <n> in JSON template
func getFirstNApps(n int) (apps []*marathon.Application) {
	content, err := readTextFile(jsonPath)
	if err != nil {
		return
	}
	group, err := parseJSON(content)
	if err != nil {
		return
	}
	return group.Apps[:n]
}

func lenTemplateApps() int {
	content, err := readTextFile(jsonPath)
	if err != nil {
		return -1
	}
	group, err := parseJSON(content)
	if err != nil {
		return -2
	}
	return len(group.Apps)
}

func groupID() string {
	content, err := readTextFile(jsonPath)
	if err != nil {
		return ""
	}
	group, err := parseJSON(content)
	if err != nil {
		return ""
	}
	return group.ID
}

func parseJSON(content []byte) (group *marathon.Group, err error) {
	group = &marathon.Group{}
	err = json.Unmarshal(content, group)
	if err != nil {
		log.Printf("unmarshal to json error: %v\n", err)
		return
	}
	return
}

func readTextFile(path string) (content []byte, err error) {
	if _, err = os.Stat(path); err != nil {
		log.Printf("stat file error: %v\n", err)
		return
	}
	content, err = ioutil.ReadFile(jsonPath)
	if err != nil {
		log.Printf("read file error: %v\n", err)
		return
	}
	return content, nil
}
