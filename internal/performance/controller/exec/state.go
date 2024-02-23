package controllerExec

var (
	runningRoom string

	suspendLog bool
)

func GetRunningRoom() string {
	return runningRoom
}
func SetRunningRoom(val string) {
	runningRoom = val
}

func IsLogSuspend() bool {
	return suspendLog
}

func SuspendLog() {
	suspendLog = true
}

func ResumeLog() {
	suspendLog = false
}
