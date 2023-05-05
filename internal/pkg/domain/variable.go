package domain

type Variable struct {
	Id                    uint   `json:"id"`
	Name                  string `json:"name"`
	Value                 string `json:"value"`
	AvailableForCurrScope bool   `json:"availableForCurrScope"`
}
