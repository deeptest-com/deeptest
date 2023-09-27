package jslibHelper

import (
	"sync"
	"time"
)

var (
	AgentLoadedLibs sync.Map
)

func GetAgentCache(id uint) (val time.Time) {
	inf, ok := AgentLoadedLibs.Load(id)

	if ok {
		val = inf.(time.Time)
	}

	return
}

func SetAgentCache(id uint, val time.Time) {
	AgentLoadedLibs.Store(id, val)
}
