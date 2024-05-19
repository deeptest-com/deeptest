package test

import (
	"log"
	"regexp"
	"testing"
)

func TestReg(t *testing.T) {
	str := "${g_var1} - {{'url is ' + escape('https://baidu.com'))}} {{1+1}}"

	reg := regexp.MustCompile(`\${(\+?[_A-Za-z][_A-Za-z0-9]*)}|(?U:{{(.+)}})`)

	arr := reg.FindAllStringSubmatch(str, -1)

	log.Println(arr)
}
