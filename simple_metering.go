package goraven

import (
	"encoding/xml"
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

func (r *Raven) SetFastPoll() {
}

func (r *Raven) GetProfileData() {
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
	XMLName             xml.Name `xml:"InstantaneousDemand"`
	DeviceMacId         string   `xml:"DeviceMacId"`
	MeterMacId          string   `xml:"MeterMacId"`
	TimeStamp           string   `xml:"TimeStamp"`
	CurrentUsage        string   `xml:"Demand"`
	Multiplier          string   `xml:"Multiplier"`
	Divisor             string   `xml:"Divisor"`
	DigitsRight         string   `xml:"DigitsRight"`
	DigitsLeft          string   `xml:"DigitsLeft"`
	SuppressLeadingZero string   `xml:"SuppressLeadingZero"`
	StartDate           string   `xml:"StartDate"`
}

// Notify: LastPeriodUsage
type LastPeriodUsage struct {
	XMLName             xml.Name `xml:"InstantaneousDemand"`
	DeviceMacId         string   `xml:"DeviceMacId"`
	MeterMacId          string   `xml:"MeterMacId"`
	TimeStamp           string   `xml:"TimeStamp"`
	LastUsage           string   `xml:"Demand"`
	Multiplier          string   `xml:"Multiplier"`
	Divisor             string   `xml:"Divisor"`
	DigitsRight         string   `xml:"DigitsRight"`
	DigitsLeft          string   `xml:"DigitsLeft"`
	SuppressLeadingZero string   `xml:"SuppressLeadingZero"`
	StartDate           string   `xml:"StartDate"`
	EndDate             string   `xml:"EndDate"`
}

