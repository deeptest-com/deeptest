package convert

type driver struct {
	FilePath string
}

func (d *driver) Data(data []byte) (res interface{}) {
	return
}

func (d *driver) setFilePath(filePath string) {
	d.FilePath = filePath
}
