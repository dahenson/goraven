package goraven

import (
	"encoding/xml"
	"math"
	"strconv"
)

// Send the GET_CURRENT_PRICE command to get the price information. Set the
// refresh element to Y to force the RAVEn to get the information from the
// meter, not the cache. The RAVEn will send a PriceCluster notification in
// response.
func (r *Raven) GetCurrentPrice() error {
	return r.simpleCommand("get_current_price", false)
}

// Not Implemented
func (r *Raven) SetCurrentPrice() {
}

// GetPrice() is a convenience function to get a correctly formatted
// floating point number of the Price contained in the PriceCluster
func (p *PriceCluster) GetPrice() (float64, error) {
	price, err := strconv.ParseInt(p.Price, 0, 0)
	if err != nil {
		return 0, err
	}

	digits, err := strconv.ParseInt(p.TrailingDigits, 0, 0)
	if err != nil {
		return 0, err
	}

	divisor := math.Pow(10, float64(digits))

	return (float64(price) / divisor), nil
}

// Notify: PriceCluster
type PriceCluster struct {
	XMLName        xml.Name `xml:"PriceCluster"`
	DeviceMacId    string   `xml:"DeviceMacId"`
	MeterMacId     string   `xml:"MeterMacId"`
	TimeStamp      string   `xml:"TimeStamp"`
	Price          string   `xml:"Price"`
	Currency       string   `xml:"Currency"`
	TrailingDigits string   `xml:"TrailingDigits"`
	Tier           string   `xml:"Tier"`
	TierLabel      string   `xml:"TierLabel,omitempty"`
	RateLabel      string   `xml:"RateLabel,omitempty"`
}
