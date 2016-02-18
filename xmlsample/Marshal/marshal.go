// XmlSample project main.go
package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Rules1 struct {
	XMLName xml.Name
}

type Rules2 struct {
	Ignore string `xml:"-"`
}

type Rules3 struct {
	ExplicitAttr string `xml:"expAttr,attr"`
}

type Rules4 struct {
	Attr string `xml:",attr"`
}

type Rules5 struct {
	CharData1 string `xml:",chardata"`
	Test      string
	CharData2 string `xml:",chardata"`
}

type Rules6 struct {
	CData string `xml:",cdata"` // 1.6版本才支持
}

type Rules7 struct {
	InnerData1 string `xml:",innerxml"`
	Test       string
	InnerData2 string `xml:",innerxml"`
}

type Rules8 struct {
	Comment1 string `xml:",comment"`
	Test     string
	Comment2 string `xml:",comment"`
}

type Rules9 struct {
	OmitEmpty1 bool `xml:",omitempty"` // 赋值false则不会进行xml序列化
	OmitEmpty2 bool `xml:",omitempty"`
}

type Rules10 struct {
	Rules10_1
}
type Rules10_1 struct {
	OuterField1 string
	OuterField2 string
}

func MarshalWithCheckError(v interface{}) []byte {
	bytes, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}

	return bytes
}
func main() {
	// 生成xml
	fmt.Println("Test Marshal Rule 1")
	rule1 := &Rules1{}
	bytes := MarshalWithCheckError(rule1)
	os.Stdout.Write(bytes)
	os.Stdout.Write([]byte("\n"))

	fmt.Println("Test Marshal Rule 2")
	rule2 := &Rules2{Ignore: "ignore"}
	bytes = MarshalWithCheckError(rule2)
	os.Stdout.Write(bytes)
	os.Stdout.Write([]byte("\n"))

	fmt.Println("Test Marshal Rule 3")
	rule3 := &Rules3{ExplicitAttr: "expattr"}
	bytes = MarshalWithCheckError(rule3)
	os.Stdout.Write(bytes)
	os.Stdout.Write([]byte("\n"))

	fmt.Println("Test Marshal Rule 4")
	rule4 := &Rules4{Attr: "attr"}
	bytes = MarshalWithCheckError(rule4)
	os.Stdout.Write(bytes)
	os.Stdout.Write([]byte("\n"))

	fmt.Println("Test Marshal Rule 5")
	rule5 := &Rules5{CharData1: "character data1", Test: "test", CharData2: "character data2"}
	bytes = MarshalWithCheckError(rule5)
	os.Stdout.Write(bytes)
	os.Stdout.Write([]byte("\n"))

	fmt.Println("Test Marshal Rule 6")
	rule6 := &Rules6{CData: "12321"}
	bytes = MarshalWithCheckError(rule6)
	os.Stdout.Write(bytes)
	os.Stdout.Write([]byte("\n"))

	fmt.Println("Test Marshal Rule 7")
	rule7 := &Rules7{InnerData1: "inner data1", Test: "test", InnerData2: "inner data2"}
	bytes = MarshalWithCheckError(rule7)
	os.Stdout.Write(bytes)
	os.Stdout.Write([]byte("\n"))

	fmt.Println("Test Marshal Rule 8")
	rule8 := &Rules8{Comment1: "Comment1", Test: "test", Comment2: "Comment2"}
	bytes = MarshalWithCheckError(rule8)
	os.Stdout.Write(bytes)
	os.Stdout.Write([]byte("\n"))

	fmt.Println("Test Marshal Rule 9")
	rule9 := &Rules9{OmitEmpty1: false, OmitEmpty2: true}
	bytes = MarshalWithCheckError(rule9)
	os.Stdout.Write(bytes)
	os.Stdout.Write([]byte("\n"))

	fmt.Println("Test Marshal Rule 10")
	rule10 := &Rules10{}
	rule10.OuterField1 = "outer field1"
	rule10.OuterField2 = "outer field2"
	bytes = MarshalWithCheckError(rule10)
	os.Stdout.Write(bytes)
	os.Stdout.Write([]byte("\n"))
}
