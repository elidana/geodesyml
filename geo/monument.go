package geo

import (
	"encoding/xml"

	"github.com/ozym/geodesyml/gmd"
	"github.com/ozym/geodesyml/gml"
)

const monumentCodeSpacePrefix = "urn:ga-gov-au:monument"

type MonumentSiteName struct {
	gml.Name
}

func NewMonumentSiteName(siteName string) MonumentSiteName {
	return MonumentSiteName{
		Name: gml.Name{
			Name: siteName,
			CodeType: gml.CodeType{
				monumentCodeSpacePrefix + "-siteName",
			},
		},
	}
}

type MonumentFourCharacterID struct {
	gml.Name
}

func NewMonumentFourCharacterID(siteName string) MonumentFourCharacterID {
	return MonumentFourCharacterID{
		Name: gml.Name{
			Name: siteName,
			CodeType: gml.CodeType{
				monumentCodeSpacePrefix + "-fourCharacterID",
			},
		},
	}
}

type MonumentIersDOMESNumber struct {
	gml.Name
}

func NewMonumentIersDOMESNumber(siteName string) MonumentIersDOMESNumber {
	return MonumentIersDOMESNumber{
		Name: gml.Name{
			Name: siteName,
			CodeType: gml.CodeType{
				monumentCodeSpacePrefix + "-iersDOMESNumber",
			},
		},
	}
}

type MonumentCdpNumber struct {
	gml.Name
}

func NewMonumentCdpNumber(siteName string) MonumentCdpNumber {
	return MonumentCdpNumber{
		Name: gml.Name{
			Name: siteName,
			CodeType: gml.CodeType{
				monumentCodeSpacePrefix + "-cdpNumber",
			},
		},
	}
}

type Status struct {
}

type Remark struct {
}

type MonumentPropertyType struct {
}

type RelativeOffsetType struct {
}

type SupplementaryMarkPropertyType struct {
	Monument
	RelativeOffsets []RelativeOffsetType `xml:"relativeOffset,omitempty"`
}

type AbstractMonument struct {
	gml.AbstractFeatureType

	Type          gml.CodeWithAuthorityType         `xml:"geo:type"`
	Status        *Status                           `xml:"geo:status,omitempty"`
	InstalledBy   *gmd.CI_ResponsibleParty_Property `xml:"gmd:installedBy"`
	installedDate *gml.TimePositionType             `xml:"gml:installedDate"`
	Remarks       []Remark

	Monument           *MonumentPropertyType           `xml:"geo:monument,omitempty"`
	SupplementaryMarks []SupplementaryMarkPropertyType `xml:"geo:SupplementaryMarkPropertyType,omitempty"`
}

type Monument struct {
	XMLName xml.Name `xml:"geo:Monument"`

	gml.Link

	AbstractMonument

	SiteName        *MonumentSiteName
	FourCharacterID *MonumentFourCharacterID
	IersDOMESNumber *MonumentIersDOMESNumber
	CdpNumber       *MonumentCdpNumber

	Inscription string `xml:"geo:inscription"`
	/*
	 <element name="monumentDescription" type="gml:CodeWithAuthorityType"/>
	 <element name="height" type="geo:SingleValueType"/>
	 <element name="foundation" type="gml:CodeWithAuthorityType"/>
	 <element name="foundationDepth" type="geo:SingleValueType"/>
	 <element name="markerDescription" type="string"/>
	 <element name="geologicCharacteristic" type="gml:CodeWithAuthorityType"/>
	 <element name="bedrockType" type="gml:CodeWithAuthorityType"/>
	 <element name="bedrockCondition" type="gml:CodeWithAuthorityType"/>
	 <element name="fractureSpacing" type="gml:CodeWithAuthorityType"/>
	 <element name="faultZonesNearby" type="gml:CodeWithAuthorityType"/>
	 <element name="distanceActivity" type="gml:CodeWithAuthorityType" minOccurs="0"/>
	*/
	FaultZonesNearby gml.CodeWithAuthorityType  `xml:"geo:faultZonesNearby"`
	DistanceActivity *gml.CodeWithAuthorityType `xml:"geo:distanceActivity,omitempty"`
}

func NewMonument() Monument {
	return Monument{
	//
	}
}

func (m *Monument) AddSiteName(siteName string) {
	m.SiteName = &MonumentSiteName{
		Name: gml.Name{
			Name: siteName,
			CodeType: gml.CodeType{
				monumentCodeSpacePrefix + "-siteName",
			},
		},
	}
}

func (m *Monument) AddFourCharacterIDName(fourCharacterID string) {
	m.FourCharacterID = &MonumentFourCharacterID{
		Name: gml.Name{
			Name: fourCharacterID,
			CodeType: gml.CodeType{
				monumentCodeSpacePrefix + "-fourCharacterID",
			},
		},
	}
}

func (m *Monument) AddIersDOMESNumber(iersDOMESNumber string) {
	m.IersDOMESNumber = &MonumentIersDOMESNumber{
		Name: gml.Name{
			Name: iersDOMESNumber,
			CodeType: gml.CodeType{
				monumentCodeSpacePrefix + "-iersDOMESNumber",
			},
		},
	}
}

func (m *Monument) AddCdpNumber(cdpNumber string) {
	m.CdpNumber = &MonumentCdpNumber{
		Name: gml.Name{
			Name: cdpNumber,
			CodeType: gml.CodeType{
				monumentCodeSpacePrefix + "-cdpNumber",
			},
		},
	}
}
