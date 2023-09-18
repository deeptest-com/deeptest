package test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestSample2(t *testing.T) {
	jsn := "[{\"title\":\"Wake up to WonderWidgets!\",\"type\":\"all\"},{\"items\":[\"Why \\u003cem\\u003eWonderWidgets\\u003c/em\\u003e are great\",\"Who \\u003cem\\u003ebuys\\u003c/em\\u003e WonderWidgets\"],\"title\":\"Overview\",\"type\":\"all\"}]"
	var obj interface{}

	err := json.Unmarshal([]byte(jsn), &obj)

	log.Println(err)
}
