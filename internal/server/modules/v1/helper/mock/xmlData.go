package mockHelper

import "encoding/xml"

var (
	str = `
		<Person>
			<FullName>Aaron</FullName>
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
type Root struct {
	XMLName xml.Name `xml:"root"`
	Res     Result   `xml:"Person"`
}
