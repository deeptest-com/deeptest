package service

import (
	"encoding/json"
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/jinzhu/copier"
)

type remote struct {
}

func (s *remote) GetUserInfoByToken(token, origin string) (user v1.UserInfo, err error) {
	url := fmt.Sprintf("%s/levault/usrsvr/Usr/GetUserInfoByToken", origin)
	body, _ := json.Marshal(map[string]string{"token": token})
	headers := make([]domain.Header, 0)
	headers = append(headers, []domain.Header{
		{
			Name:  "Tenant-Id",
			Value: "1632931640315338752", //TODO 做saas后可以考虑提到配置文件里
		},
		{
			Name:  "Ds-Tenant-Id",
			Value: "0",
		},
		{
			Name:  "Token",
			Value: token,
		},
		{
			Name:  "lang",
			Value: "zh_cn",
		}}...,
	)
	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers:  &headers,
		Body:     string(body),
	}

	var resp domain.DebugResponse
	resp, err = httpHelper.Post(httpReq)
	if err != nil {
		logUtils.Infof("meta get method detail failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("meta get method detail failed, response %v", resp)
		err = fmt.Errorf("meta get method detail failed, response %v", resp)
		return
	}

	respContent := struct {
		Code int
		Data v1.OtherUserInfo
		Msg  string
	}{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	copier.CopyWithOption(&user, respContent.Data, copier.Option{DeepCopy: true})

	return
}
