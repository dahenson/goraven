package goraven

func (r *Raven) GetInstantaneousDemand() {
}

func (r *Raven) GetCurrentSummationDelivered() {
}

// Get the accumulated usage information from the RAVEn. The RAVEn will send
// a CurrentPeriodUsage notification in response. Note that this command will
// not cause the current period consumption total to be updated. To do this,
// send a GetCurrentSummationDelivered command with Refresh set to Y.
func (r *Raven) GetCurrentPeriodUsage() error {
	return r.simpleCommand("get_current_period_usage")
}

// Get the previous period accumulation data from the RAVEn. The RAVEn will
// send a LastPeriodUsage notification in response.
func (r *Raven) GetLastPeriodUsage() error {
	return r.simpleCommand("get_last_period_usage")
}

// Force the RAVEn to roll over the current period to the last period and
// initialize the current period.
func (r *Raven) CloseCurrentPeriod() error {
	return r.simpleCommand("close_current_period")
}

func (r *Raven) SetFastPoll() {
}

func (r *Raven) GetProfileData() {
}
