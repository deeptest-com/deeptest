package exec

import (
	"fmt"
)

const (
	VarNameVuCount = "VarNameVuCount"
)

func IncreaseVuCount(room, serverAddress string) (ret int) {
	remoteValue := AddRemoteVal(room, getVuCountKey(), serverAddress)

	ret = int(remoteValue.Value)

	return
}

func ResetVuCount(room, serverAddress string) (ret int) {
	ResetRemoteVal(room, getVuCountKey(), serverAddress)

	remoteValue := GetRemoteVal(room, getVuCountKey(), serverAddress)

	ret = int(remoteValue.Value)

	return
}

func getVuCountKey() string {
	return fmt.Sprintf("%s", VarNameVuCount)
}
