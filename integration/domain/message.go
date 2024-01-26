package integrationDomain

type ApprovalReq struct {
	CreatorId    string   `json:"creatorId"`    //批发起人username 必需
	ApproveIds   []string `json:"approveIds"`   //json数组，审批人username 必需
	ApprovalType int      `json:"approvalType"` //审批类型 1或 2并，目前只支持1
	CcIds        []string `json:"ccIds"`        //抄送人username

	Title     string `json:"title"`     //审批标题 必需
	Content   string `json:"content"`   //童批内容 必需
	SourceIds []int  `json:"sourceIds"` //必需 来源id 关联表id，通常是主键，如果没有设置一个[0]
	Remark    string `json:"remark"`    //申请原因
	NotifyUrl string `json:"notifyUrl"` //通知地址
}
