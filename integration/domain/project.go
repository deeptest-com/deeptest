package integrationDomain

type ProjectReq struct {
	Products    []uint   `json:"products"`
	Spaces      []string `json:"spaces"`
	SyncMembers bool     `json:"syncMembers"`
}

type SpaceRole struct {
	Id        uint   `json:"id"`
	RoleName  string `json:"roleName"`
	RoleValue string `json:"roleValue"`
	Remark    string `json:"remark"`
}

type ProductItem struct {
	ProductBaseItem
	Children []*ProductItem `json:"children"`
}

type ProductBaseItem struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type SpaceItem struct {
	Name        string `json:"name"`
	NameEngAbbr string `json:"nameEngAbbr"`
}

type ProjectDetail struct {
	Products []ProductBaseItem `json:"products"`
	Spaces   []SpaceItem       `json:"spaces"`
}
