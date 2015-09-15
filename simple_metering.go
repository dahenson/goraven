package goraven

import (
	"encoding/xml"
	"strconv"
)

// Get the demand information from the RAVEn. If refresh is true, the device
// gets the information from the meter instead of from its cache.
func (r *Raven) GetInstantaneousDemand(refresh bool) error {
	return r.simpleCommand("get_instantaneous_demand", refresh)
}

// Get the summation information from the RAVEn. If refresh is true, the device
// gets the information from the meter instead of from its cache.
func (r *Raven) GetCurrentSummationDelivered(refresh bool) error {
	return r.simpleCommand("get_current_summation_delivered", refresh)
}

// Get the accumulated usage information from the RAVEn. The RAVEn will send
// a CurrentPeriodUsage notification in response. Note that this command will
// not cause the current period consumption total to be updated. To do this,
// send a GetCurrentSummationDelivered command with Refresh set to Y.
func (r *Raven) GetCurrentPeriodUsage() error {
	return r.simpleCommand("get_current_period_usage", false)
}

// Get the previous period accumulation data from the RAVEn. The RAVEn will
// send a LastPeriodUsage notification in response.
func (r *Raven) GetLastPeriodUsage() error {
	return r.simpleCommand("get_last_period_usage", false)
}

// Force the RAVEn to roll over the current period to the last period and
// initialize the current period.
func (r *Raven) CloseCurrentPeriod() error {
	return r.simpleCommand("close_current_period", false)
}

// Not Implemented
func (r *Raven) SetFastPoll() {
}

// Not Implemented
func (r *Raven) GetProfileData() {
}

// GetDemand() is a convenience function that returns a correctly formatted
// floating point number for the Demand field
func (i *InstantaneousDemand) GetDemand() (float64, error) {
	return getFloat64(i.Demand, i.Multiplier, i.Divisor)
}

// GetSummationDelivered() is a convenience function that returns a correctly
// formatted floating point number for the Current Summation Delivered field
func (c *CurrentSummationDelivered) GetSummationDelivered() (float64, error) {
	return getFloat64(c.SummationDelivered, c.Multiplier, c.Divisor)
}

// GetSummationReceived() is a convenience function that returns a correctly
// formatted floating point number for the Current Summation Received field
func (c *CurrentSummationDelivered) GetSummationReceived() (float64, error) {
	return getFloat64(c.SummationReceived, c.Multiplier, c.Divisor)
}

func getFloat64(dem, mult, div string) (float64, error) {
	i, err := strconv.ParseInt(dem, 0, 0)
	demand := float64(i)
	if err != nil {
		return 0.0, err
	}

	i, err = strconv.ParseInt(mult, 0, 0)
	multiplier := float64(i)
	if multiplier == 0 {
		multiplier = 1
	}
	if err != nil {
		return 0.0, err
	}

	i, err = strconv.ParseInt(div, 0, 0)
	divisor := float64(i)
	if divisor == 0 {
		divisor = 1
	}
	if err != nil {
		return 0.0, err
	}

	return (demand * multiplier / divisor), nil
}

// Notify: InstantaneousDemand
type InstantaneousDemand struct {
	XMLName             xml.Name `xml:"InstantaneousDemand"`
	DeviceMacId         string   `xml:"DeviceMacId"`
	MeterMacId          string   `xml:"MeterMacId"`
	TimeStamp           string   `xml:"TimeStamp"`
	Demand              string   `xml:"Demand"`
	Multiplier          string   `xml:"Multiplier"`
	Divisor             string   `xml:"Divisor"`
	DigitsRight         string   `xml:"DigitsRight"`
	DigitsLeft          string   `xml:"DigitsLeft"`
	SuppressLeadingZero string   `xml:"SuppressLeadingZero"`
}

// Notify: CurrentSummationDelivered
type CurrentSummationDelivered struct {
	XMLName             xml.Name `xml:"CurrentSummationDelivered"`
	DeviceMacId         string   `xml:"DeviceMacId"`
	MeterMacId          string   `xml:"MeterMacId"`
	TimeStamp           string   `xml:"TimeStamp"`
	SummationDelivered  string   `xml:"SummationDelivered"`
	SummationReceived   string   `xml:"SummationReceived"`
	Multiplier          string   `xml:"Multiplier"`
	Divisor             string   `xml:"Divisor"`
	DigitsRight         string   `xml:"DigitsRight"`
	DigitsLeft          string   `xml:"DigitsLeft"`
	SuppressLeadingZero string   `xml:"SuppressLeadingZero"`
}

// Notify: CurrentPeriodUsage
type CurrentPeriodUsage struct {
	XMLName             xml.Name `xml:"CurrentPeriodUsage"`
	DeviceMacId         string   `xml:"DeviceMacId"`
	MeterMacId          string   `xml:"MeterMacId"`
	TimeStamp           string   `xml:"TimeStamp"`
	CurrentUsage        string   `xml:"CurrentUsage"`
	Multiplier          string   `xml:"Multiplier"`
	Divisor             string   `xml:"Divisor"`
	DigitsRight         string   `xml:"DigitsRight"`
	DigitsLeft          string   `xml:"DigitsLeft"`
	SuppressLeadingZero string   `xml:"SuppressLeadingZero"`
	StartDate           string   `xml:"StartDate"`
}

// Notify: LastPeriodUsage
type LastPeriodUsage struct {
	XMLName             xml.Name `xml:"LastPeriodUsage"`
	DeviceMacId         string   `xml:"DeviceMacId"`
	MeterMacId          string   `xml:"MeterMacId"`
	TimeStamp           string   `xml:"TimeStamp"`
	LastUsage           string   `xml:"LastUsage"`
	Multiplier          string   `xml:"Multiplier"`
	Divisor             string   `xml:"Divisor"`
	DigitsRight         string   `xml:"DigitsRight"`
	DigitsLeft          string   `xml:"DigitsLeft"`
	SuppressLeadingZero string   `xml:"SuppressLeadingZero"`
	StartDate           string   `xml:"StartDate"`
	EndDate             string   `xml:"EndDate"`
}

