package consts

import "fmt"

type Integers []int

func (n Integers) ToString() (ret string) {
	for _, number := range n {
		if ret == "" {
			ret += fmt.Sprintf("%v", number)
		} else {
			ret += fmt.Sprintf(",%v", number)
		}
	}

	return ret
}

type TenantId string
