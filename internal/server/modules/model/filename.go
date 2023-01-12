package model

type FileName struct {
	BaseModel
	FileName string `json:"file_name"`
	PathName string `json:"path_name"`
}

func NewFileName(filename string, pathname string) *FileName {
	return &FileName{
		FileName: filename,
		PathName: pathname,
	}

}

func (FileName) TableName() string {
	return "biz_filename"
}
