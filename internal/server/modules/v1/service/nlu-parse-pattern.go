package service

import (
	"github.com/aaronchen2k/deeptest/internal/comm/consts"
	_logUtils "github.com/aaronchen2k/deeptest/internal/pkg/lib/log"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"regexp"
	"time"
)

type NluParsePatternService struct {
	NluPatternService *NluPatternService `inject:""`
}

func NewNluParsePatternService() *NluParsePatternService {
	return &NluParsePatternService{}
}

func (s *NluParsePatternService) Parse(projectId uint, req serverDomain.NluReq) (ret serverDomain.NluResp) {
	ret.Code = -1
	nluResult := serverDomain.NluResult{
		Intent: &serverDomain.Intent{
			Confidence: 1,
		},
		StartTime: time.Now(),
	}

	text := req.Text
	if len(serverConsts.NluPatterns) == 0 {
		s.NluPatternService.Reload()
	}
	tasks := serverConsts.NluPatterns

	found := false

OuterLoop:
	for _, task := range tasks {
		for _, intent := range task.Intents {
			for _, sent := range intent.Sents {
				regx := regexp.MustCompile("(?i)" + sent.Example)

				slotArr := regx.FindStringSubmatch(text)

				if slotArr != nil { // matched
					found = true

					nluResult.Intent.ID = int64(intent.Id)
					nluResult.Intent.Name = intent.Name
					nluResult.IntentRanking = append(nluResult.IntentRanking, serverDomain.IntentRanking{
						Name:       intent.Name,
						Confidence: 1,
					})
					nluResult.Text = text
					nluResult.Intent.Sent = serverDomain.Sent{
						ID:   int64(sent.Id),
						Name: sent.Example,
					}

					s.popEntities(slotArr, text, sent, &nluResult)

					break OuterLoop
				}
			}
		}
	}

	nluResult.EndTime = time.Now()
	if !found {
		nluResult.Intent.Confidence = 0
		nluResult.Entities = make([]serverDomain.Entity, 0)
	}

	ret.SetResult(nluResult)
	ret.Code = 1

	return
}

func (s *NluParsePatternService) popEntities(slotArr []string, text string,
	sent serverDomain.NluSent, resp *serverDomain.NluResult) {

	if len(slotArr) <= len(sent.Slots) {
		_logUtils.Errorf("error to parse %s with pattern %s", text, sent.Example)
	}

	for index, item := range slotArr {
		sent.Slots[index] = slotArr[index]

		entity := serverDomain.Entity{Extractor: consts.Pattern, ConfidenceEntity: 1}
		entity.Value = item

		resp.Entities = append(resp.Entities, entity)

		index += 2
	}
}
