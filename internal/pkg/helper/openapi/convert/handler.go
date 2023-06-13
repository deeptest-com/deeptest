package convert

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
)

type IDriver interface {
	toOpenapi() *openapi3.T
	Doc(data []byte)
}

type Handler struct {
	driver IDriver
}

func NewHandler(driverType DriverType, data []byte) *Handler {
	h := &Handler{}
	h.driver = initDriver(driverType, data)
	return h
}

func (h *Handler) ToOpenapi() *openapi3.T {
	return h.driver.toOpenapi()
}

func initDriver(driverType DriverType, data []byte) (driver IDriver) {

	switch driverType {
	case POSTMAN:
		driver = newPostman()
	case YAPI:
		driver = newYApi()
	case SWAGGER:
		driver = newSwagger()
	default:
		panic(fmt.Errorf("dirver error"))
	}

	driver.Doc(data)
	return
}
