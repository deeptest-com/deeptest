package agentExec

type ProcessorPerformanceRunner struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntityBase

	Ip       string `json:"ip"`
	WebPort  uint   `json:"webPort"`
	GrpcPort uint   `json:"grpcPort"`
	Weight   uint   `json:"weight"`
}

func (entity ProcessorPerformanceRunner) Run(processor *Processor, session *ExecSession) (err error) {
	return
}
