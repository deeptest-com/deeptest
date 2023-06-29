package agentExec

import "os"

var (
	DemoTestSite = GetDemoSite()
)

func GetDemoSite() (ret string) {
	ret = os.Getenv("DEMO_TEST_SITE")

	return
}
