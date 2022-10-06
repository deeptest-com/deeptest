package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
)

type ProcessorInterface struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntity

	Method            consts.HttpMethod      `gorm:"default:GET" json:"method"`
	Url               string                 `json:"url"`
	Params            []domain.Param         `gorm:"-" json:"params"`
	Headers           []domain.Header        `gorm:"-" json:"headers"`
	Body              string                 `gorm:"default:{}" json:"body"`
	BodyType          consts.HttpContentType `gorm:"default:json" json:"bodyType"`
	AuthorizationType consts.AuthorType      `gorm:"default:''" json:"authorizationType"`
	PreRequestScript  string                 `gorm:"default:''" json:"preRequestScript"`
	ValidationScript  string                 `gorm:"default:''" json:"validationScript"`

	BasicAuth   domain.BasicAuth   `gorm:"-" json:"basicAuth"`
	BearerToken domain.BearerToken `gorm:"-" json:"bearerToken"`
	OAuth20     domain.OAuth20     `gorm:"-" json:"oauth20"`
	ApiKey      domain.ApiKey      `gorm:"-" json:"apiKey"`

	InterfaceID uint `json:"interfaceID"`
	ProjectID   uint `json:"projectID"`
}

func (p ProcessorInterface) Run(s *Session) (log Result, err error) {
	logUtils.Infof("interface entity")

	//variableMap, _ := GetVariablesByInterfaceAndProcessor(p.ProcessorID, p.ProjectID)
	//requestHelper.ReplaceAll(&req, variableMap)
	//
	//resp, err := s.InterfaceService.Test(req)
	//if err != nil {
	//	return
	//}
	//
	//logPo, err := s.ExecLogService.CreateInterfaceLog(req, resp, parentLog)
	//if err != nil {
	//	return
	//}
	//
	//logExtractors, err := s.ExtractorService.ExtractInterface(interf, resp, &logPo)
	//logCheckpoints, status, err := s.CheckpointService.CheckInterface(interf, resp, &logPo)
	//
	//// send msg to client
	//reqContent, _ := json.Marshal(req)
	//respContent, _ := json.Marshal(resp)
	//
	//interfaceLog := &domain.ExecLog{
	//	Id:                logPo.ID,
	//	Name:              interfaceProcessor.Name,
	//	ProcessorCategory: consts.ProcessorInterface,
	//	ProcessorType:     consts.ProcessorInterfaceDefault,
	//	ParentId:          parentLog.PersistentId,
	//
	//	InterfaceId:  interf.ID,
	//	ReqContent:   string(reqContent),
	//	RespContent:  string(respContent),
	//	ResultStatus: status,
	//
	//	InterfaceExtractorsResult:  logExtractors,
	//	InterfaceCheckpointsResult: logCheckpoints,
	//}
	//
	//*parentLog.Logs = append(*parentLog.Logs, interfaceLog)
	//execHelper.SendExecMsg(*interfaceLog, wsMsg)

	log = Result{
		Name:        p.Name,
		InterfaceId: p.InterfaceID,
		//ReqContent:   string(reqContent),
		//RespContent:  string(respContent),
		//ResultStatus: status,
		//
		//InterfaceExtractorsResult:  logExtractors,
		//InterfaceCheckpointsResult: logCheckpoints,
	}

	return
}
