package k6Comm

type ExecutorsType string

const (
	SharedIterations     ExecutorsType = "shared-iterations"
	PerVUIterations      ExecutorsType = "per-vu-iterations"
	ConstantVUs          ExecutorsType = "constant-vus"
	RampingVUs           ExecutorsType = "ramping-vus"
	ConstantArrivalRate  ExecutorsType = "constant-arrival-rate"
	RampingArrivalRate   ExecutorsType = "ramping-arrival-rate"
	ExternallyControlled ExecutorsType = "externally-controlled"
)

func (e ExecutorsType) String() string {
	return string(e)
}
