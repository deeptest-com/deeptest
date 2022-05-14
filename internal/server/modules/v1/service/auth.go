package service

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
)

type AuthService struct {
	InvocationRepo *repo.InvocationRepo `inject:""`
	InterfaceRepo  *repo.InterfaceRepo  `inject:""`
}

func (s AuthService) GenOAuth2AccessToken(req model.InterfaceOAuth20) (err error) {
	s.InterfaceRepo.UpdateOAuth20(req.InterfaceId, req)

	return
}
