package convert

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
)

type IDriver interface {
	toOpenapi() (*openapi3.T, error)
	Doc(data []byte)
	setFilePath(filePath string)
}

type Handler struct {
	driver IDriver
}

func NewHandler(driverType DriverType, data []byte, filePath string) *Handler {
	h := &Handler{}
	h.driver = initDriver(driverType, data, filePath)
	return h
}

func (h *Handler) ToOpenapi() (*openapi3.T, error) {
	return h.driver.toOpenapi()
}

func initDriver(driverType DriverType, data []byte, filePath string) (driver IDriver) {

	switch driverType {
	case POSTMAN:
		driver = newPostman()
	case YAPI:
		driver = newYApi()
	case SWAGGER2:
		driver = newSwaggerV2()
	case SWAGGER3:
		driver = newSwaggerV3()
	default:
		panic(fmt.Errorf("dirver error"))
	}

	driver.Doc(data)
	driver.setFilePath(filePath)
	return
}
