package grpcHelper

import (
	"context"
	"errors"
	"fmt"
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	grpcCore "github.com/aaronchen2k/deeptest/internal/pkg/helper/grpc/core"
	grpcDomain "github.com/aaronchen2k/deeptest/internal/pkg/helper/grpc/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

var (
	reGetFuncArg = regexp.MustCompile("\\( (.*) \\) returns")
)

// Handler hold all handler methods
type Handler struct {
	G *grpcCore.GrpCox
}

// InitHandler Constructor
func InitHandler() *Handler {
	return &Handler{
		G: grpcCore.InitGrpCox(),
	}
}

func (h *Handler) List(req serverDomain.GrpcReq) (services, methods []string, err error) {
	address := req.Address
	service := req.Service
	method := req.Method
	useTls := req.UseTls
	isRestartConn := req.RestartConn

	metaData := []string{}
	for _, item := range req.MetaData {
		metaData = append(metaData, fmt.Sprintf("%s:%s", item.Key, item.Value))
	}

	h.G.SetReflectHeaders(metaData...)

	res, err := h.G.GetResource(context.Background(), address, !useTls, isRestartConn)
	if err != nil {
		return
	}

	if service == "" { // parse, load services
		services, err = res.List("")
		if err != nil {
			logUtils.Infof("grpc List error: %s", err.Error())
			return
		}

	} else if service != "" && method == "" { // select service, load methods
		methods, err = res.List(service)
		if err != nil {
			logUtils.Infof("grpc List error: %s", err.Error())
			return
		}

	} else if service != "" && method != "" { // reload, load both services and methods
		services, err = res.List("")
		if err != nil {
			logUtils.Infof("grpc List error: %s", err.Error())
			return
		}

		methods, err = res.List(service)
		if err != nil {
			logUtils.Infof("grpc List error: %s", err.Error())
			return
		}
	}

	h.G.Extend(address)

	return
}

// ListWithProto handling client request for service list with proto
func (h *Handler) ListWithProto(req serverDomain.GrpcReq) (services, methods []string, err error) {
	address := req.Address
	service := req.Service
	method := req.Method
	useTls := req.UseTls
	isRestartConn := req.RestartConn
	protoName := req.ProtoName
	protoPath := req.ProtoPath

	// convert uploaded file to list of Proto struct
	protos := []grpcCore.Proto{}

	fileData, err := os.Open(protoPath)
	if err != nil {
		return
	}
	defer fileData.Close()

	content, err1 := io.ReadAll(fileData)
	if err1 != nil {
		err = err1
		return
	}

	protos = append(protos, grpcCore.Proto{
		Name:    protoName,
		Content: content,
	})

	res, err := h.G.GetResourceWithProto(context.Background(), address, !useTls, isRestartConn, protos)
	if err != nil {
		return
	}

	if service == "" { // parse, load services
		services, err = res.List("")
		if err != nil {
			logUtils.Infof("grpc List error: %s", err.Error())
			return
		}

	} else if service != "" && method == "" { // select service, load methods
		methods, err = res.List(service)
		if err != nil {
			logUtils.Infof("grpc List error: %s", err.Error())
			return
		}

	} else if service != "" && method != "" { // reload, load both services and methods
		services, err = res.List("")
		if err != nil {
			logUtils.Infof("grpc List error: %s", err.Error())
			return
		}

		methods, err = res.List(service)
		if err != nil {
			logUtils.Infof("grpc List error: %s", err.Error())
			return
		}
	}

	h.G.Extend(address)

	return
}

func (h *Handler) DescribeFunc(req serverDomain.GrpcReq) (ret grpcDomain.Desc, err error) {
	address := req.Address
	funcName := req.Method
	useTls := req.UseTls
	isRestartConn := req.RestartConn

	res, err := h.G.GetResource(context.Background(), address, !useTls, isRestartConn)
	if err != nil {
		return
	}

	// get param
	result, _, isClientStreaming, isServerStreaming, err := res.Describe(funcName, "func")
	if err != nil {
		return
	}
	matchedParams := reGetFuncArg.FindStringSubmatch(result)
	if len(matchedParams) < 2 {
		errors.New("Invalid Func Type")
		return
	}

	// describe func
	result, template, _, _, err := res.Describe(matchedParams[1], "param")
	if err != nil {
		return
	}

	h.G.Extend(address)

	ret = grpcDomain.Desc{
		Schema:            result,
		Template:          template,
		IsClientStreaming: isClientStreaming,
		IsServerStreaming: isServerStreaming,
	}

	return
}

func (h *Handler) InvokeFunc(req serverDomain.GrpcReq) (ret grpcDomain.InvRes, err error) {
	address := req.Address
	funcName := req.Method
	useTls := req.UseTls
	isRestartConn := req.RestartConn
	message := req.Message
	isClientStreaming := req.IsClientStreaming
	isServerStreaming := req.IsServerStreaming

	res, err := h.G.GetResource(context.Background(), address, !useTls, isRestartConn)
	if err != nil {
		return
	}

	metaData := []string{}
	for _, item := range req.MetaData {
		metaData = append(metaData, fmt.Sprintf("%s:%s", item.Key, item.Value))
	}

	// get param
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(600)*time.Second)
	body := strings.NewReader(message)

	isClientStreaming = true
	result, timer, err := res.Invoke(ctx, metaData, funcName, body, isClientStreaming, isServerStreaming)
	if err != nil {
		return
	}

	h.G.Extend(address)

	ret = grpcDomain.InvRes{
		Time:   timer.String(),
		Result: result,
	}

	return
}
