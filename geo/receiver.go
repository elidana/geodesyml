package geo

import (
	"encoding/xml"

	"github.com/ozym/geodesyml/gml"
)

type Receiver struct {
	XMLName xml.Name `xml:"geo:gnssReceiver"`
	gml.Id
	//string   `xml:"gml:id,attr"`

	ManufacturerSerialNumber string  `xml:"geo:manufacturerSerialNumber"`
	ReceiverType             string  `xml:"geo:receiverType"`
	SatelliteSystem          string  `xml:"geo:satelliteSystem"`
	SerialNumber             string  `xml:"geo:serialNumber"`
	FirmwareVersion          string  `xml:"geo:firmwareVersion"`
	ElevationCutoffSetting   float32 `xml:"geo:elevationCutoffSetting"`
	DateInstalled            string  `xml:"geo:dateInstalled"`
	DateRemoved              string  `xml:"geo:dateRemoved"`
	TemperatureStabilization string  `xml:"geo:temperatureStabilization"`
	Notes                    string  `xml:"geo:notes"`
}
