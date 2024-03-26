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
	Products    []ProductBaseItem `json:"products"`
	Spaces      []SpaceItem       `json:"spaces"`
	Engineering []EngineeringItem `json:"engineeringItem"`
}

type UserInfo struct {
	Username string `json:"username"`
	WxName   string `json:"wxName"`
	RealName string `json:"realName"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

type ProjectInfo struct {
	Name        string     `json:"name"`        // 名称
	NameEngAbbr string     `json:"nameEngAbbr"` // 英文名称缩写
	SpaceAdmins []UserInfo `json:"spaceAdmins"` // 空间管理员
}

type UserMenuPermission struct {
	Permission string               `json:"permission"`
	Children   []UserMenuPermission `json:"children"`
}

type SpaceMembersAndRolesItem struct {
	ProjectEngAbbr string         `json:"ProjectEngAbbr"`
	UserBaseInfo   []UserRoleInfo `json:"UserBaseInfo"`
}

type UserRoleInfo struct {
	UserInfo
	Role []struct {
		Id    uint   `json:"Id"`
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"role"`
	RoleValues []string `json:"roleValues"`
}
