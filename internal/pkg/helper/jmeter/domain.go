package jmeterHelper

import (
	"encoding/xml"
)

type JmeterTestPlan struct {
	XMLName    xml.Name `xml:"jmeterTestPlan"`
	Version    string   `xml:"version,attr"`
	Properties string   `xml:"properties,attr"`
	JMeter     string   `xml:"jmeter,attr"`
	HashTree   HashTree `xml:"hashTree"`
}

type HashTree struct {
	XMLName xml.Name      `xml:"hashTree"`
	Items   []interface{} `xml:"items"`
}

type TestPlan struct {
	XMLName    xml.Name     `xml:"TestPlan"`
	Name       string       `xml:"testname,attr"`
	GUIClass   string       `xml:"guiclass,attr"`
	TestClass  string       `xml:"testclass,attr"`
	Enabled    bool         `xml:"enabled,attr"`
	StringProp []StringProp `xml:"stringProp"`
}

type ThreadGroup struct {
	XMLName     xml.Name     `xml:"ThreadGroup"`
	GUIClass    string       `xml:"guiclass,attr"`
	TestClass   string       `xml:"testclass,attr"`
	Enabled     bool         `xml:"enabled,attr"`
	Name        string       `xml:"testname,attr"`
	StringProp  []StringProp `xml:"stringProp"`
	ElementProp ElementProp  `xml:"elementProp"`
}

type HTTPSamplerProxy struct {
	XMLName     xml.Name      `xml:"HTTPSamplerProxy"`
	Name        string        `xml:"testname,attr"`
	GUIClass    string        `xml:"guiclass,attr"`
	TestClass   string        `xml:"testclass,attr"`
	Enabled     bool          `xml:"enabled,attr"`
	StringProp  []StringProp  `xml:"stringProp"`
	BoolProp    []BoolProp    `xml:"boolProp"`
	ElementProp []ElementProp `xml:"elementProp"`
}

type ResultCollector struct {
	XMLName   xml.Name `xml:"ResultCollector"`
	Enabled   bool     `xml:"enabled,attr"`
	GUIClass  string   `xml:"guiclass,attr"`
	TestClass string   `xml:"testclass,attr"`
	Name      string   `xml:"testname,attr"`
}

type ElementProp struct {
	Name           string           `xml:"name,attr"`
	Type           string           `xml:"elementType,attr"`
	GUIClass       string           `xml:"guiclass,attr"`
	TestClass      string           `xml:"testclass,attr"`
	Enabled        bool             `xml:"enabled,attr"`
	StringProp     []StringProp     `xml:"stringProp"`
	BoolProp       []BoolProp       `xml:"boolProp"`
	CollectionProp []CollectionProp `xml:"collectionProp"`
}

type CollectionProp struct {
	Name        string        `xml:"name,attr"`
	ElementProp []ElementProp `xml:"elementProp"`
}

type StringProp struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

type BoolProp struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}
