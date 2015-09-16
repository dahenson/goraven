package goraven

import (
	"encoding/xml"
	"strconv"
	"time"
)

// Send the GET_TIME command to get the current time. The RAVEn will send
// a TimeCluster notification in response
func (r *Raven) GetTime() error {
	return r.simpleCommand("get_time", false)
}

// GetUTCTime() is a convenience function to get the UTC time contained
// in a TimeCluster notify
func (t *TimeCluster) GetUTCTime() (time.Time, error) {
	return getTime(t.UTCTime)
}

// GetLocalTime() is a convenience function to get the Local time contained
// in a TimeCluster notify
func (t *TimeCluster) GetLocalTime() (time.Time, error) {
	return getTime(t.LocalTime)
}

func getTime(hex string) (time.Time, error) {
	sec, err := strconv.ParseInt(hex, 0, 0)
	if err != nil {
		return time.Unix(0, 0), err
	}

	return time.Unix(sec+946684800, 0), nil
}

type TimeCluster struct {
	XMLName     xml.Name `xml:"TimeCluster"`
	DeviceMacId string   `xml:"DeviceMacId"`
	MeterMacId  string   `xml:"MeterMacId"`
	UTCTime     string   `xml:"UTCTime"`
	LocalTime   string   `xml:"LocalTime"`
}
