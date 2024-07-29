package service

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	grpcHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/grpc"
	grpcDomain "github.com/aaronchen2k/deeptest/internal/pkg/helper/grpc/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"github.com/jinzhu/copier"
	"path/filepath"
	"strings"
	"sync"
)

var (
	handleStore sync.Map
)

type GrpcInterfaceService struct {
	GrpcInterfaceRepo     *repo.GrpcInterfaceRepo     `inject:""`
	DiagnoseInterfaceRepo *repo.DiagnoseInterfaceRepo `inject:""`
}

func (s *GrpcInterfaceService) GetDebugData(tenantId consts.TenantId, id int) (ret domain.GrpcDebugData, err error) {
	diagnose, err := s.DiagnoseInterfaceRepo.Get(tenantId, uint(id))

	copier.CopyWithOption(&ret, diagnose, copier.Option{
		DeepCopy: true,
	})

	po, err := s.GrpcInterfaceRepo.Get(tenantId, diagnose.DebugInterfaceId)
	if err != nil {
		return
	}

	copier.CopyWithOption(&ret, po, copier.Option{
		DeepCopy: true,
	})

	return
}

func (s *GrpcInterfaceService) SaveDebugData(data domain.GrpcDebugData, tenantId consts.TenantId) (err error) {
	po := model.GrpcInterface{}
	copier.CopyWithOption(&po, data, copier.Option{DeepCopy: true})

	err = s.GrpcInterfaceRepo.SaveDebugData(po, tenantId)

	return
}

func (s *GrpcInterfaceService) ParseProto(req serverDomain.GrpcReq) (
	ret serverDomain.GrpcParseResp, err error) {

	handler := s.getHandler(req.ExecUuid)
	if handler == nil {
		return
	}

	if req.ProtoPath != "" {
		req.ProtoPath = filepath.Join(consts.WorkDir, req.ProtoPath)
	}

	var services, methods []string

	if req.ProtoSrc == "local" {
		services, methods, err = handler.ListWithProto(req)
	} else {
		services, methods, err = handler.List(req)
	}
	if err != nil {
		return
	}

	s.popServices(&ret, services)
	s.popMethods(&ret, methods)

	return
}

func (s *GrpcInterfaceService) DescribeFunction(req serverDomain.GrpcReq) (
	ret grpcDomain.Desc, err error) {

	handler := s.getHandler(req.ExecUuid)
	if handler == nil {
		return
	}

	ret, _ = handler.DescribeFunc(req)

	ret.Schema = ret.Schema
	ret.Template = _stringUtils.FormatJsonStr(ret.Template)

	return
}

func (s *GrpcInterfaceService) InvokeFunc(req serverDomain.GrpcReq) (
	ret grpcDomain.InvRes, err error) {

	handler := s.getHandler(req.ExecUuid)
	if handler == nil {
		return
	}

	ret, err = handler.InvokeFunc(req)
	if err != nil {
		return
	}

	return
}

func (s *GrpcInterfaceService) ListActiveConn(req serverDomain.GrpcReq) (result []string, err error) {
	handler := s.getHandler(req.ExecUuid)
	if handler == nil {
		return
	}

	result, err = handler.G.ListActiveConn(strings.TrimSpace(req.Address))
	if err != nil {
		return
	}

	return
}

func (s *GrpcInterfaceService) DeleteHandle(req serverDomain.GrpcReq) (err error) {
	handler := s.getHandler(req.ExecUuid)
	if handler == nil {
		return
	}

	err = handler.G.CloseActiveConns("all")
	if err != nil {
		return
	}

	handleStore.Delete(req.ExecUuid)

	return
}

func (s *GrpcInterfaceService) popServices(resp *serverDomain.GrpcParseResp, result []string) (
	ret grpcDomain.Desc, err error) {

	for _, item := range result {
		if strings.Contains(item, "grpc.reflection.") {
			continue
		}

		service := serverDomain.GrpcService{
			Name: strings.TrimSpace(item),
		}
		(*resp).Services = append((*resp).Services, service)
	}

	return
}

func (s *GrpcInterfaceService) popMethods(resp *serverDomain.GrpcParseResp, result []string) (
	ret grpcDomain.Desc, err error) {

	for _, item := range result {
		service := serverDomain.GrpcMethod{
			Name: strings.TrimSpace(item),
		}
		(*resp).Methods = append((*resp).Methods, service)
	}

	return
}

func (s *GrpcInterfaceService) getHandler(execUuid string) (ret *grpcHelper.Handler) {
	obj, ok := handleStore.Load(execUuid)
	if !ok {
		inst := grpcHelper.InitHandler()
		handleStore.Store(execUuid, inst)

		obj, ok = handleStore.Load(execUuid)
		if !ok {
			return
		}
	}

	ret = obj.(*grpcHelper.Handler)

	return
}
