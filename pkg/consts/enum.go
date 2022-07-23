package _consts

type ResultCode int

const (
	ResultCodeSuccess ResultCode = 0
	ResultCodeFail    ResultCode = 1
)

func (e ResultCode) Int() int {
	return int(e)
}

type WsMsgCategory string

const (
	Output WsMsgCategory = "output"

	Run    WsMsgCategory = "run"
	Result WsMsgCategory = "result"
	Error  WsMsgCategory = "error"

	Communication WsMsgCategory = "communication"
	Unknown       WsMsgCategory = ""
)

func (e WsMsgCategory) String() string {
	return string(e)
}
