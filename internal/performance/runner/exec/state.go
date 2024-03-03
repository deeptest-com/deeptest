package runnerExec

var (
	isRunnerRunning = false
)

func IsRunnerTestRunning() bool {
	return isRunnerRunning
}

func SetRunnerTestRunning(val bool) {
	isRunnerRunning = val
}
