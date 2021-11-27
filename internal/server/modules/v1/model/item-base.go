package model

type BaseItem struct {
	// build in
	Title    string `json:"title" gorm:"comment:'主题'"`
	Priority string `json:"priority" gorm:"comment:'优先级'"`
	Assignee uint   `json:"assignee" gorm:"comment:'处理人'"`
	Creator  uint   `json:"creator" gorm:"comment:'创建人'"`
	Reporter uint   `json:"reporter" gorm:"comment:'提交人'"` // 创建人可指定提交人为他人
	Tester   uint   `json:"tester" gorm:"comment:'测试人员'"`
	Desc     string `json:"desc" gorm:"column:descr;comment:'描述'"`
	Version  string `json:"version" gorm:"comment:'所属版本'"`
	Parent   uint   `json:"parent" gorm:"comment:'所属父工作项'"`

	Attachments []Attachment `json:"attachments" gorm:"-;comment:'附件'"`
	Labels      []Label      `json:"labels" gorm:"-;comment:'标签'"`
	Relations   []BaseItem   `json:"relations" gorm:"-;comment:'关联工作项'"`
}
