package mockHelper

import (
	"encoding/xml"
)

var (
	str = `
		<Person>
			<FullName>Aaron Chen</FullName>
			<Company>Example Inc.</Company>
			<Email where="home">
				<Addr>462826@qq.com</Addr>
			</Email>
			<Email where='work'>
				<Addr>master@deeptest.com</Addr>
			</Email>
			<Group>
				<Value>admin</Value>
				<Value>dev</Value>
			</Group>
			<City>Suzhou</City>
			<State>Jiangsu</State>
		</Person>
	`
)

func GetXmlData() (result Result) {
	result = Result{Name: "none", Phone: "none"}

	xml.Unmarshal([]byte(str), &result)

	return
}

func FormatXml(str string) (ret string) {
	x := node{}
	_ = xml.Unmarshal([]byte(str), &x)
	buf, _ := xml.MarshalIndent(x, "", "\t")

	ret = string(buf)
	return
}

type node struct {
	Attr     []xml.Attr
	XMLName  xml.Name
	Children []node `xml:",any"`
	Text     string `xml:",chardata"`
}

type Email struct {
	Where string `xml:"where,attr"`
	Addr  string
}
type Address struct {
	City, State string
}
type Result struct {
	XMLName xml.Name `xml:"Person"`
	Name    string   `xml:"FullName"`
	Phone   string
	Email   []Email
	Groups  []string `xml:"Group>Value"`
	Address
}
