package gml

const gmlNameSpace = "http://www.opengis.net/gml/3.2"

type GmlNameSpace struct {
	NameSpace string `xml:"xmlns:gml,attr"`
}

func NewGmlNameSpace() GmlNameSpace {
	return GmlNameSpace{
		NameSpace: gmlNameSpace,
	}
}

type Id struct {
	Id string `xml:"gml:id,attr,omitempty"`
}

func NewId(s string) Id {
	return Id{
		Id: s,
	}
}

type Link struct {
	Link string `xml:"xlink:href,attr,omitempty"`
}

func NewLink(s string) Link {
	return Link{
		Link: s,
	}
}
