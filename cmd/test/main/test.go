package main

import (
	builtin "github.com/aaronchen2k/deeptest/internal/pkg/buildin"
	"log"
)

func main() {
	str := builtin.MD5("value")

	log.Println(str)
}
