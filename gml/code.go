package gml

import (
//	"encoding/xml"
)

/*
type CodeType struct {
	CodeSpace string `xml:"codespace,attr,omitempty"`
}
*/

type CodeWithAuthorityType struct {
	CodeSpace string `xml:"codespace,attr,omitempty"`
}

/*
type Name struct {
	XMLName xml.Name `xml:"gml:name"`

	CodeType
	Name string `xml:",chardata"`
}

func NewName(codeSpace, name string) Name {
	return Name{
		CodeType: CodeType{
			codeSpace,
		},
		Name: name,
	}
}
*/
