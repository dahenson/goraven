package goraven

// Send the GET_TIME command to get the current time. The RAVEn will send
// a TimeCluster notification in response
func (r *Raven) GetTime() {
	return r.simpleCommand("get_time", false)
}

type TimeCluster struct {
	XMLName     xml.Name `xml:"TimeCluster"`
	DeviceMacId string   `xml:"DeviceMacId"`
	MeterMacId  string   `xml:"MeterMacId"`
	UTCTime     string   `xml:"UTCTime"`
	LocalTime   string   `xml:"LocalTime"`
}
