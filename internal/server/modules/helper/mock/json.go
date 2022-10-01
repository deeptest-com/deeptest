package mockHelper

import (
	"encoding/json"
)

var (
	JsonStr = `{
		"name": "John",
		"age"      : 26,
		"address"  : {
		  "streetAddress": "naist street",
		  "city"         : "Nara",
		  "postalCode"   : "630-0192"
		},
		"phoneNumbers": [
		  {
			"type"  : "cellphone",
			"number": "13912345678"
		  },
		  {
			"type"  : "telephone",
			"number": "0512-65911234"
		  }
		]
	}`
)

func GetJsonData() (person Person) {
	json.Unmarshal([]byte(JsonStr), &person)
	return
}

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address struct {
		StreetAddress string `json:"streetAddress"`
		City          string `json:"city"`
		PostalCode    string `json:"postalCode"`
	} `json:"address"`
	PhoneNumbers []struct {
		Type   string `json:"type"`
		Number string `json:"number"`
	} `json:"phoneNumbers"`
}
