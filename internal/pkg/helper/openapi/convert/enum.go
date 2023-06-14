package convert

type DriverType string

const (
	POSTMAN DriverType = "postman"
	YAPI    DriverType = "yapi"
	SWAGGER DriverType = "swagger"
)

type DataSyncType string

const (
	FullCover DataSyncType = "full_cover" //完全覆盖
	CopyAdd   DataSyncType = "copy_add"   //复制新增
)

func (e DataSyncType) String() string {
	return string(e)
}
