package _stringUtils

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

func Uuid() string {
	uid := uuid.NewV4().String()
	return strings.Replace(uid, "-", "", -1)
}

func UuidWithSep() string {
	uid := uuid.NewV4().String()
	return uid
}
