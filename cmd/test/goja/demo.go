package main

import (
	"fmt"

	"github.com/dlclark/regexp2"
)

// demo   regex2
func demo() {
	regex := regexp2.MustCompile(`((?<=\d{7})\w+?(?=\d))`, regexp2.None)
	match, err := regex.FindStringMatch("1111111wwww111123a")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(match.GroupByNumber(1).String())
	}
}
