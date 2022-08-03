package domain

type Variable struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Value      string `json:"value"`
	Expression string `json:"expression"`
}
