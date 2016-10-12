package autoscaling

import (
	"log"

	"github.com/LinkerNetworks/gwMonitor/conf"
)

func scalePgwUp() {
	group, err := getPgwGroup()
	if err != nil {
		log.Printf("get pgw group error: %v\n", err)
		return
	}

	scaleto := len(group.Apps) + conf.OptionsReady.PgwScaleStep

	n := lenTemplateApps()
	if scaleto > n {
		log.Println("all template apps has started up, wont scale up")
		return
	}

	group.Apps = getFirstNApps(scaleto)

	err = updateGroup(group)
	if err != nil {
		log.Printf("update pgw group error: %v\n", err)
	}
}

func scaleSgwUp() {
	group, err := getSgwGroup()
	if err != nil {
		log.Printf("get sgw group error: %v\n", err)
		return
	}

	scaleto := len(group.Apps) + conf.OptionsReady.SgwScaleStep

	n := lenTemplateApps()
	if scaleto > n {
		log.Println("all template apps has started up, wont scale up")
		return
	}

	group.Apps = getFirstNApps(scaleto)

	err = updateGroup(group)
	if err != nil {
		log.Printf("update sgw group error: %v\n", err)
	}
}
