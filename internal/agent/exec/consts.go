package agentExec

import (
	"log"
	"os"
)

const (
	ServerApiPath = "api/v1"
)

var (
	DemoTestSite = GetDemoSite()
)

func GetDemoSite() (ret string) {
	ret = os.Getenv("DemoTestSite")
	log.Printf("DemoTestSite: %s\n", ret)

	return
}
