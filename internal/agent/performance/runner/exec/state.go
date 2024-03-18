package runnerExec

var (
	isRunnerRunning = false
)

func IsTestRunning() bool {
	return isRunnerRunning
}

func SetTestRunning(val bool) {
	isRunnerRunning = val
}
