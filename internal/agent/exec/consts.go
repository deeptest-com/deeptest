package agentExec

import (
	"log"
	"os"
)

var (
	DemoTestSite = GetDemoSite()
)

func GetDemoSite() (ret string) {
	ret = os.Getenv("DemoTestSite")
	log.Printf("DemoTestSite: %s\n", ret)

	return
}
