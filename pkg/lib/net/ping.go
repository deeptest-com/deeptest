package _netUtils

import (
	"fmt"
	"net"
	"time"
)

func Ping(host string, port string) (ret bool) {
	address := fmt.Sprintf("%s:%s", host, port)

	conn, err := net.DialTimeout("tcp", address, time.Second*2)
	if err != nil {
		return false
	}

	defer conn.Close()

	return true
}
