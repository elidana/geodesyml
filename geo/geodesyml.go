package geo

import (
	"encoding/xml"

	"github.com/ozym/geodesyml/gml"
)

/*
const NameSpace = "urn:xml-gov-au:icsm:egeodesy:0.2"

type GeoNameSpace struct {
	MyNameSpace string `xml:"xmlns:geo,attr"`
}

func NewGeoNameSpace() GeoNameSpace {
	return GeoNameSpace{
		MyNameSpace: NameSpace,
	}
}
*/

type GeodesyML struct {
	XMLName xml.Name `xml:"geo:GeodesyML"`

	// Name Spaces
	GeoNameSpace
	gml.GmlNameSpace

	Monument Monument

	Receivers []Receiver `xml:",omitempty"`
}

func NewGeodesyML() *GeodesyML {

	return &GeodesyML{
		GeoNameSpace: NewGeoNameSpace(),
		GmlNameSpace: gml.NewGmlNameSpace(),
	}
}

func (x *GeodesyML) AddMonument(m Monument) {
	x.Monument = m
}

func (x *GeodesyML) AddReceiver(r Receiver) {
	x.Receivers = append(x.Receivers, r)
}
func (x *GeodesyML) AddReceivers(rs []Receiver) {
	for _, r := range rs {
		x.AddReceiver(r)
	}
}

// shouldn't be here
func (x GeodesyML) Marshal() ([]byte, error) {
	h := xml.Header
	s, err := xml.MarshalIndent(x, "  ", "    ")
	if err != nil {
		return nil, err
	}
	return append([]byte(h), append(s, '\n')...), nil
}
