package goraven

// Get the meter information. The RAVEn will send a MeterInfo notification in
// response.
func (r *Raven) GetMeterInfo() error {
	return r.simpleCommand("get_meter_info")
}

// Get the status of the device on the network. The RAVEn will send a
// NetworkInfo notification in response.
func (r *Raven) GetNetworkInfo() error {
	return r.simpleCommand("get_network_info")
}

func (r *Raven) SetMeterInfo() {
}

