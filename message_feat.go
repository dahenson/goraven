package goraven

import (
	"encoding/xml"
)

// GetMessage gets the current message. The RAVEn will send a MessageCluster
// notification in response.
func (r *Raven) GetMessage(refresh bool) error {
	return r.simpleCommand("get_message", refresh)
}

// Notify: MessageCluster
type MessageCluster struct {
	XMLName              xml.Name `xml:"MessageCluster"`
	DeviceMacId          string   `xml:"DeviceMacId"`
	MeterMacId           string   `xml:"MeterMacId"`
	TimeStamp            string   `xml:"TimeStamp"`
	Id                   string   `xml:"Id"`
	Text                 string   `xml:"Text"`
	ConfirmationRequired string   `xml:"ConfirmationRequired"`
	Queued               string   `xml:"Queued"`
}
