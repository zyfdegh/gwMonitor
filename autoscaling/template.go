package autoscaling

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	marathon "github.com/gambol99/go-marathon"
)

const (
	// JSONPath is Marathon JSON template file path
	JSONPath = "pgw.json"
)

// return apps from 0 to n in JSON template
func getFirstNApps(n int) (apps []*marathon.Application) {
	content, err := readTextFile(JSONPath)
	if err != nil {
		return
	}
	group, err := parseJSON(content)
	if err != nil {
		return
	}
	return group.Apps[:n]
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
	content, err = ioutil.ReadFile(JSONPath)
	if err != nil {
		log.Printf("read file error: %v\n", err)
		return
	}
	return content, nil
}
