package serverDomain

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type DiagnoseCurlLoadReq struct {
	EndpointInterfaceId uint `json:"endpointInterfaceId"`
	CaseId              uint `json:"caseId"`
	DiagnoseId          uint `json:"diagnoseId"`

	EnvironmentId uint          `json:"environmentId"`
	ProjectId     int           `json:"projectId"`
	UserId        uint          `json:"userId"`
	UsedBy        consts.UsedBy `json:"usedBy"`
}
