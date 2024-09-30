package serverDomain

import "github.com/deeptest-com/deeptest/internal/pkg/consts"

type DiagnoseCurlLoadReq struct {
	// from endpoint interface list
	EndpointId      uint              `json:"endpointId"`
	InterfaceMethod consts.HttpMethod `json:"interfaceMethod"`

	// from debug page
	DebugInterfaceId    uint `json:"debugInterfaceId"`
	EndpointInterfaceId uint `json:"endpointInterfaceId"`
	CaseId              uint `json:"caseId"`
	DiagnoseId          uint `json:"diagnoseId"`

	EnvironmentId uint          `json:"environmentId"`
	ProjectId     int           `json:"projectId"`
	UserId        uint          `json:"userId"`
	UsedBy        consts.UsedBy `json:"usedBy"`
	FromDefine    bool          `json:"fromDefine"`
}
