package ptutils

import "strings"

func GetGrpcAddress(serverAddress string) (ret string) {
	ret = strings.ReplaceAll(serverAddress, "8086", "9528")

	return
}
