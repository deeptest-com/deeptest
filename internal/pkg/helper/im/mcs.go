package im

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
)

type Mcs struct {
	ServiceType consts.MessageServiceType `json:"serviceType"`
	Data        interface{}               `json:"data"`
}

type mcsReq struct {
	McsAppid    string      `json:"mcsAppid"`
	Tool        int         `json:"tool"`
	ServiceType int         `json:"serviceType"`
	Data        interface{} `json:"data"`
}

type mcsInfoRes struct {
	mcsCommonRes
	Data mcsInfoResData `json:"data"`
}

type mcsApprovalRes struct {
	mcsCommonRes
	Data mcsApprovalResData `json:"data"`
}

type mcsInfoResData struct {
	MsgId string `json:"msgId"`
}

type mcsApprovalResData struct {
	InstanceId string `json:"instanceId"`
}

type mcsCommonRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type EnterpriseWechatInfoData struct {
	Content     string   `json:"content"`     //消息内容
	UserAccount []string `json:"userAccount"` //字符串数组 接收人列表，由第三方用户唯一识别标志组成，如企微用户账号
	ImAppid     int      `json:"imAppid"`     //企微应用id
	ChatId      string   `json:"chatId"`      //群聊id
	MsgType     int      `json:"msgType"`     //消息类型：0:markdown消息 1:图文消息，默认markdown消息
}

type EnterpriseWechatApprovalData struct {
	CreatorId    string         `json:"creatorId"`    //批发起人第三方账号
	ApproveIds   []string       `json:"approveIds"`   //json数组，审批人第三方账号
	ApprovalType int            `json:"approvalType"` //审批类型 1或 2并，目前只支持或签1
	CcIds        []string       `json:"ccIds"`        //抄送人账号
	TemplateId   string         `json:"templateId"`   //审批模板id
	ButtonDetail []ButtonDetail `json:"buttonDetail"`
	SummaryList  []string       `json:"summaryList"` //摘要
	NotifyUrl    string         `json:"notifyUrl"`   //接收审批结果url
}

type ButtonDetail struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	Data string `json:"data"`
}

type EnterpriseWechatInfoContent struct {
	Articles []EnterpriseWechatInfoContentArticles `json:"articles"`
}

type EnterpriseWechatInfoContentArticles struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Picurl      string `json:"picurl"`
	Appid       string `json:"appid"`
	Pagepath    string `json:"pagepath"`
}

func (s *Mcs) SendMessage() (msgId string, err error) {
	url := fmt.Sprintf("%s/api/v1/mcsCall/serviceRequest", config.CONFIG.Mcs.Url)
	req := mcsReq{
		McsAppid:    config.CONFIG.Mcs.McsAppid,
		Tool:        1, // TODO 1企微 2飞书 3钉钉 如果以后接入了飞书/钉钉，就配置到server.yaml中
		ServiceType: s.transferToMcsServiceType(),
		Data:        s.Data,
	}

	body, err := json.Marshal(req)
	if err != nil {
		logUtils.Infof("marshal mcs request data failed, error, %s", err.Error())
		return
	}

	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Body:     string(body),
	}

	resp, err := httpHelper.Post(httpReq)
	if err != nil {
		logUtils.Infof("send message by mcs failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK {
		logUtils.Infof("send message by mcs failed, response %v", resp)
		return
	}

	if s.ServiceType == consts.ServiceTypeApproval {
		respContent := mcsApprovalRes{}
		err = json.Unmarshal([]byte(resp.Content), &respContent)
		if err != nil {
			logUtils.Infof(err.Error())
		}
		msgId = respContent.Data.InstanceId
	} else {
		respContent := mcsInfoRes{}
		err = json.Unmarshal([]byte(resp.Content), &respContent)
		if err != nil {
			logUtils.Infof(err.Error())
		}
		msgId = respContent.Data.MsgId
	}

	return
}

func (s *Mcs) transferToMcsServiceType() (serviceType int) {
	if s.ServiceType == consts.ServiceTypeApproval {
		serviceType = 1
	} else if s.ServiceType == consts.ServiceTypeInfo {
		serviceType = 2
	}

	return
}
