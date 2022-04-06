package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//
// microsft .res to android string.xml conversion
//

//
// http://cuto.unirita.co.jp/gostudy/post/standard-library-xml/
// https://qiita.com/chanmitsu55/items/8268f559efa694bd1cfd
// https://qiita.com/ono_matope/items/70080cc33b75152c5c2a
// https://mike-neck.hatenadiary.com/entry/golang-xml-parse
//

type Root struct {
	XMLName xml.Name `xml:"root"`
	Text    string   `xml:",chardata"`
	Schema  struct {
		Text   string `xml:",chardata"`
		ID     string `xml:"id,attr"`
		Xmlns  string `xml:"xmlns,attr"`
		Xsd    string `xml:"xsd,attr"`
		Msdata string `xml:"msdata,attr"`
		Import struct {
			Text      string `xml:",chardata"`
			Namespace string `xml:"namespace,attr"`
		} `xml:"import"`
		Element struct {
			Text        string `xml:",chardata"`
			Name        string `xml:"name,attr"`
			IsDataSet   string `xml:"IsDataSet,attr"`
			ComplexType struct {
				Text   string `xml:",chardata"`
				Choice struct {
					Text      string `xml:",chardata"`
					MaxOccurs string `xml:"maxOccurs,attr"`
					Element   []struct {
						Text        string `xml:",chardata"`
						Name        string `xml:"name,attr"`
						ComplexType struct {
							Text     string `xml:",chardata"`
							Sequence struct {
								Text    string `xml:",chardata"`
								Element []struct {
									Text      string `xml:",chardata"`
									Name      string `xml:"name,attr"`
									Type      string `xml:"type,attr"`
									MinOccurs string `xml:"minOccurs,attr"`
									Ordinal   string `xml:"Ordinal,attr"`
								} `xml:"element"`
							} `xml:"sequence"`
							Attribute []struct {
								Text    string `xml:",chardata"`
								Name    string `xml:"name,attr"`
								Use     string `xml:"use,attr"`
								Type    string `xml:"type,attr"`
								Ref     string `xml:"ref,attr"`
								Ordinal string `xml:"Ordinal,attr"`
							} `xml:"attribute"`
						} `xml:"complexType"`
					} `xml:"element"`
				} `xml:"choice"`
			} `xml:"complexType"`
		} `xml:"element"`
	} `xml:"schema"`
	Resheader []struct {
		Text  string `xml:",chardata"`
		Name  string `xml:"name,attr"`
		Value string `xml:"value"`
	} `xml:"resheader"`
	Data []struct {
		Text    string `xml:",chardata"`
		Name    string `xml:"name,attr"`
		Space   string `xml:"space,attr"`
		Value   string `xml:"value"`
		Comment string `xml:"comment"`
	} `xml:"data"`
}

type androidString struct {
	name  string
	value int
}

// main
func main() {

	if len(os.Args) == 1 {
		log.Fatal("no args")
	}

	xmlFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer xmlFile.Close()

	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		log.Fatal(err)
	}

	var r Root

	if err := xml.Unmarshal([]byte(xmlData), &r); err != nil {
		fmt.Println("XML Unmarshal error:", err)
		return
	}

	for i := 0; i < len(r.Data); i++ {
		fmt.Print("<string name=\"")
		fmt.Print(r.Data[i].Name)
		fmt.Print(`">`)
		fmt.Print(r.Data[i].Value)
		fmt.Println("</string>")
	}

	/*
		for i, x := Range r.Data {

			fmt.Println(i)
			fmt.Println(x);

		}*/

	//fmt.Println("%v", r.Data)
	// fmt.Println(r.Datas[0].Value)
}
