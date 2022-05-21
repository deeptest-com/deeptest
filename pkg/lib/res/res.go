package _resUtils

import (
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"github.com/aaronchen2k/deeptest/res"
	"io/ioutil"
)

func ReadRes(path string) (ret []byte, err error) {
	isRelease := _commUtils.IsRelease()
	if isRelease {
		ret, err = res.Asset(path)
	} else {
		ret, err = ioutil.ReadFile(path)
	}

	return
}
