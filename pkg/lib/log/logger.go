package _logUtils

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

// init in other places for server and agent
