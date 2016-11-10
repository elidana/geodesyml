package geo

const geoNameSpace = "urn:xml-gov-au:icsm:egeodesy:0.2"

type GeoNameSpace struct {
	NameSpace string `xml:"xmlns:geo,attr"`
}

func NewGeoNameSpace() GeoNameSpace {
	return GeoNameSpace{
		NameSpace: geoNameSpace,
	}
}
