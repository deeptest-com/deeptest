package serverDomain

type LoginByOauthReq struct {
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
}

type LoginByOauthRes struct {
	ThirdPartyCommonRes
	Data LoginByOauthResData `json:"data"`
}

type LoginByOauthResData struct {
	Code     string `json:"code"`
	TenantId string `json:"tenantId"`
}

type GetTokenFromCodeReq struct {
	Code string `json:"code"`
}

type GetTokenFromCodeRes struct {
	ThirdPartyCommonRes
	Data GetTokenFromCodeResData `json:"data"`
}

type GetTokenFromCodeResData struct {
	Token      string `json:"token"`
	FreshToken string `json:"freshToken"`
	ExprieIn   string `json:"expriein"`
	CreateAt   string `json:"createAt"`
}

type FindClassByServiceCodeReq struct {
	ServiceCode string `json:"serviceCode"`
}

type FindClassByServiceCodeRes struct {
	ThirdPartyCommonRes
	Data []FindClassByServiceCodeResData `json:"data"`
}

type FindClassByServiceCodeResData struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type GetFunctionsByClassReq struct {
	ServiceCode string `json:"serviceCode"`
	ClassCode   string `json:"classCode"`
}

type GetFunctionsByClassRes struct {
	ThirdPartyCommonRes
	Data []GetFunctionsByClassResData `json:"data"`
}

type GetFunctionsByClassResData struct {
	Code        string `json:"code"`
	MessageType int    `json:"messageType"` // 0：内部方法，不能被前端调用 1：外部方法，可以被前端调用
}

type MetaGetMethodDetailReq struct {
	ClassName   string `json:"className"`
	Method      string `json:"method"`
	IncludeSelf bool   `json:"includeSelf"`
}

type MetaGetMethodDetailRes struct {
	ThirdPartyCommonRes
	Data MetaGetMethodDetailResData `json:"data"`
}

type MetaGetMethodDetailResData struct {
	Code            string `json:"code"`
	ServiceCode     string `json:"serviceCode"`
	ClassCode       string `json:"classCode"`
	RequestType     string `json:"requestType"`   //JSON/FORM
	RequestMethod   string `json:"requestMethod"` //POST
	RequestFormBody string `json:"requestFormBody"`
	RequestBody     string `json:"requestBody"`
	ResponseType    string `json:"responseType"` //JSON
	ResponseBody    string `json:"responseBody"`
}

type ThirdPartyCommonRes struct {
	Mfail  string      `json:"mfail"`
	Msg    string      `json:"msg"`
	Errors interface{} `json:"errors"`
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
