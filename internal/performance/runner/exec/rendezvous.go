package exec

import (
	"fmt"
)

const (
	VarNameRendezvousArrived = "VarNameRendezvousArrived"
	VarNameRendezvousPassed  = "VarNameRendezvousPassed"
)

func IsRendezvousReady(room, name, serverAddress string, target int) (ret int, ready bool) {
	if target == 0 {
		ready = true
		return
	}

	value := GetRemoteVal(room, getArrivedKey(name), serverAddress)

	if value == nil {
		return
	}

	ret = int(value.Value)

	if ret >= target {
		ready = true
	}

	return
}

func IncreaseRendezvousArrived(room, name, serverAddress string) (ret int) {
	remoteValue := IncreaseRemoteVal(room, getArrivedKey(name), serverAddress)

	ret = int(remoteValue.Value)

	return
}

func IncreaseRendezvousPassed(room, name, serverAddress string) (ret int) {
	remoteValue := IncreaseRemoteVal(room, getPassedKey(name), serverAddress)

	ret = int(remoteValue.Value)

	return
}

func ResetRendezvous(room, name, serverAddress string) (ret int) {
	ResetRemoteVal(room, getArrivedKey(name), serverAddress)
	ResetRemoteVal(room, getPassedKey(name), serverAddress)

	remoteValue := GetRemoteVal(room, getArrivedKey(name), serverAddress)

	ret = int(remoteValue.Value)

	return
}

func getArrivedKey(name string) string {
	return fmt.Sprintf("%s_%s", VarNameRendezvousArrived, name)
}

func getPassedKey(name string) string {
	return fmt.Sprintf("%s_%s", VarNameRendezvousPassed, name)
}
