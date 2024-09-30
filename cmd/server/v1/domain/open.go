package serverDomain

import "github.com/deeptest-com/deeptest/internal/pkg/consts"

type SaveSpaceRelatedProjectsReq struct {
	SpaceCode         string   `json:"spaceCode"`
	ProjectShortNames []string `json:"projectShortNames"`
}

type ProjectListWithRole struct {
	ID        uint            `json:"id"`
	Name      string          `json:"name"`
	ShortName string          `json:"shortName"`
	RoleName  consts.RoleType `json:"roleName"`
}
