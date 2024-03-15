package consts

type IntegrationFuncExtendStatus string

const (
	IntegrationFuncIsExtend    IntegrationFuncExtendStatus = "YES"
	IntegrationFuncIsNotExtend IntegrationFuncExtendStatus = "NO"
)

func (e IntegrationFuncExtendStatus) String() string {
	return string(e)
}

type IntegrationFuncOverridable string

const (
	IntegrationFuncCanOverridable    IntegrationFuncOverridable = "YES"
	IntegrationFuncCanNotOverridable IntegrationFuncOverridable = "NO"
)

func (e IntegrationFuncOverridable) String() string {
	return string(e)
}

type CronLecangMessageType string

const (
	CronLecangMessageTypeInner   CronLecangMessageType = "inner"
	CronLecangMessageTypeOutside CronLecangMessageType = "outside"
)

func (e CronLecangMessageType) String() string {
	return string(e)
}

type CronLecangIsExtendOverride string

const (
	CronLecangExtendOverride CronLecangIsExtendOverride = "extend_override"     //继承并重写
	CronLecangExtend         CronLecangIsExtendOverride = "extend_not_override" //继承未重写
	CronLecangNotExtend      CronLecangIsExtendOverride = "not_extend"          //未继承，自身
)

func (e CronLecangIsExtendOverride) String() string {
	return string(e)
}

type CronLecangOverridable string

const (
	CronLecangIsOverridable  CronLecangOverridable = "YES"
	CronLecangNotOverridable CronLecangOverridable = "NO"
)

func (e CronLecangOverridable) String() string {
	return string(e)
}
