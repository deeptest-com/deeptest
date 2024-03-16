package ptdomain

type ErrorAlreadyRunning struct{}

func (e *ErrorAlreadyRunning) Error() string {
	return "already_running"
}
