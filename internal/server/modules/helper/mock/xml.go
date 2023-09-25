package mockHelper

import (
	"encoding/xml"
)

var (
	XmlStr = `
		<Person>
			<FullName>Aaron Chen</FullName>
			<Email where="home">
				<Address>462826@qq.com</Address>
			</Email>
			<Email where='work'>
				<Address>master@deeptest.com</Address>
			</Email>
			<Group>
				<Sample>admin</Sample>
				<Sample>dev</Sample>
			</Group>
			<City>Suzhou</City>
			<State>Jiangsu</State>
		</Person>
	`
)

func GetXmlData() (result Result) {
	result = Result{}
	xml.Unmarshal([]byte(XmlStr), &result)

	return
}

type Result struct {
	XMLName xml.Name `xml:"Person"`
	Name    string   `xml:"FullName"`
	Phone   string
	Email   []Email
	Groups  []string `xml:"Group>Sample"`
	Address
}
type Email struct {
	Where   string `xml:"where,attr"`
	Address string
}
type Address struct {
	City, State string
}
