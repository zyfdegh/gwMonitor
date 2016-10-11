package autoscaling

import (
	"encoding/json"
	"fmt"
	"os"
)

//for debugging
func printPretty(v interface{}, mark string) (err error) {
	fmt.Printf("*********%s\n", mark)
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return
	}
	data = append(data, '\n')
	os.Stdout.Write(data)
	return
}
