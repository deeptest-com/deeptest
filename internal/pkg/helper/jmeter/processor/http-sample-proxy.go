package jmeterProcessor

import (
	"fmt"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/beevik/etree"
)

func HTTPSamplerProxy(elem *etree.Element) (processor *agentExec.Processor) {
	name := elem.SelectAttrValue("testname", "")

	entity := agentExec.ProcessorInterface{
		ProcessorEntityBase: agentExec.ProcessorEntityBase{
			Name: name,
		},
	}

	processor = &agentExec.Processor{
		ProcessorBase: agentExec.ProcessorBase{
			Name: name,
		},
		Entity: entity,
	}

	argsElems := elem.FindElements("//elementProp[@name='HTTPsampler.Arguments']/collectionProp[@name='Arguments.arguments']/elementProp")

	for _, arg := range argsElems {
		nameElem := arg.FindElement("//stringProp[@name='Argument.name']")
		valElem := arg.FindElement("//[@name='Argument.value']")

		name := GetAttrContent(nameElem)
		val := GetAttrContent(valElem)

		param := domain.Param{
			Name:  name,
			Value: val,
		}
		entity.QueryParams = append(entity.QueryParams, param)
	}

	protocolElem := elem.FindElement("[@name='HTTPSampler.protocol']")
	domainElem := elem.FindElement("[@name='HTTPSampler.domain']")
	pathElem := elem.FindElement("[@name='HTTPSampler.path']")
	methodElem := elem.FindElement("[@name='HTTPSampler.method']")

	entity.Url = fmt.Sprintf("%s://%s/%s",
		GetAttrContent(protocolElem),
		GetAttrContent(domainElem), GetAttrContent(pathElem))
	entity.Method = consts.HttpMethod(GetAttrContent(methodElem))

	elem.Child = []etree.Token{}

	return
}
