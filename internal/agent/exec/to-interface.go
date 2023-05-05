package agentExec

import "github.com/aaronchen2k/deeptest/internal/pkg/domain"

type InterfaceExecObj struct {
	DebugData domain.DebugData `json:"debugData"`

	ExecScene domain.ExecScene `json:"execScene"`
}
