// XmlSample project main.go
package main

import (
	"encoding/xml"
	"fmt"
)

type Rules1 struct {
	InnerElem string `xml:",innerxml"`
}

type Rules2 struct {
	XMLName xml.Name
}

type Rules3 struct {
	XMLName xml.Name `xml:"Test1"` // xml元素名必须匹配``描述的Test1，否则返回错误
}

type Rules4 struct {
	Attr         string `xml:",attr"`        // xml元素的属性名必须匹配结构体的Attr属性名
	ExplicitAttr string `xml:"expAttr,attr"` // xml元素的属性名必须匹配``描述的expAttr名字
}

type Rules5 struct {
	CharData1 []byte `xml:",chardata"` // xml元素包含的字符内容
	CharData2 string `xml:",chardata"` // xml元素包含的字符内容
}

type Rules6 struct {
	Comment1 string `xml:",comment"`
}

type Rules7 struct {
	SubElem1 string `xml:"SubElem1"`
	SubElem2 string `xml:">SubSubElem"`
}

type Rules8 struct {
	Sub Rules8_1
}

type Rules8_1 struct {
	XMLName   xml.Name `xml:"SubElem"`
	InnerElem string   `xml:",innerxml"`
}

type Rules9 struct {
	WithoutFlag string
}

type Rules10 struct {
	AnyMatch string `xml:",any"`
}

type Rules11 struct {
	Rules11_1
}

type Rules11_1 struct {
	OuterField1 string
	OuterField2 string
}

type Rules12 struct {
	Ignore    string `xml:"-"`
	NotIgnore string
}

func UnmarshalWithCheckError(bytes []byte, v interface{}) {
	err := xml.Unmarshal(bytes, v)
	if err != nil {
		panic(err)
	}
}

func main() {
	// 解析xml
	fmt.Println("Test Unmarshal Rule 1")
	data := `<Test>test</Test>`
	var rule1 Rules1
	UnmarshalWithCheckError([]byte(data), &rule1)
	fmt.Printf("%s\n", rule1.InnerElem)

	data = `<Test></Test>`
	fmt.Println("Test Unmarshal Rule 2")
	var rule2 Rules2
	UnmarshalWithCheckError([]byte(data), &rule2)
	fmt.Printf("%s\n", rule2.XMLName)

	data = `<Test1></Test1>`
	fmt.Println("Test Unmarshal Rule 3")
	var rule3 Rules3
	UnmarshalWithCheckError([]byte(data), &rule3)
	fmt.Printf("%s\n", rule3.XMLName)

	data = `<Test Attr="attr" expAttr="expattr"></Test>`
	fmt.Println("Test Unmarshal Rule 4")
	var rule4 Rules4
	UnmarshalWithCheckError([]byte(data), &rule4)
	fmt.Printf("%s %s\n", rule4.Attr, rule4.ExplicitAttr)

	data = `<Test>chardata<TestElem></TestElem>chardata1</Test>`
	fmt.Println("Test Unmarshal Rule 5")
	var rule5 Rules5
	rule5.CharData2 = "test rule 5"
	UnmarshalWithCheckError([]byte(data), &rule5)
	fmt.Printf("%s : %s\n", rule5.CharData1, rule5.CharData2)

	data = `<Test><!--ABOUT--></Test>`
	fmt.Println("Test Unmarshal Rule 6")
	var rule6 Rules6
	UnmarshalWithCheckError([]byte(data), &rule6)
	fmt.Printf("%s\n", rule6.Comment1)

	data = `<Test><SubElem1>subelem1</SubElem1><SubElem2><SubSubElem>subsubelem</SubSubElem></SubElem2></Test>`
	fmt.Println("Test Unmarshal Rule 7")
	var rule7 Rules7
	UnmarshalWithCheckError([]byte(data), &rule7)
	fmt.Printf("%s : %s\n", rule7.SubElem1, rule7.SubElem2)

	data = `<Test><SubElem>subinner</SubElem></Test>`
	fmt.Println("Test Unmarshal Rule 8")
	var rule8 Rules8
	UnmarshalWithCheckError([]byte(data), &rule8)
	fmt.Printf("%s : %s\n", rule8.Sub.XMLName, rule8.Sub.InnerElem)

	data = `<Test><WithoutFlag>without flag</WithoutFlag></Test>`
	fmt.Println("Test Unmarshal ule 9")
	var rule9 Rules9
	UnmarshalWithCheckError([]byte(data), &rule9)
	fmt.Printf("%s\n", rule9.WithoutFlag)

	data = `<Test><AnyMatchRule>any match</AnyMatchRule></Test>`
	fmt.Println("Test Unmarshal Rule 10")
	var rule10 Rules10
	UnmarshalWithCheckError([]byte(data), &rule10)
	fmt.Printf("%s\n", rule10.AnyMatch)

	data = `<Test><OuterField1>outer struct field1</OuterField1><OuterField2>outer struct field2</OuterField2></Test>`
	fmt.Println("Test Unmarshal Rule 11")
	var rule11 Rules11
	UnmarshalWithCheckError([]byte(data), &rule11)
	fmt.Printf("%s : %s\n", rule11.OuterField1, rule11.OuterField2)
	fmt.Printf("%s : %s\n", rule11.Rules11_1.OuterField1, rule11.Rules11_1.OuterField2)

	data = `<Test><Ignore>ignore</Ignore><NotIgnore>not ignore</NotIgnore></Test>`
	fmt.Println("Test Unmarshal Rule 12")
	var rule12 Rules12
	rule12.Ignore = "rules12 ignore"
	rule12.NotIgnore = "rules12 not ignore"
	UnmarshalWithCheckError([]byte(data), &rule12)
	fmt.Printf("%s : %s\n", rule12.Ignore, rule12.NotIgnore)
}
