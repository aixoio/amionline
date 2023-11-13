package data

type EventData struct {
	Success bool `json:"success"`
	Time_ms float64 `json:"time_ms"`
	Target_ip string `json:"target_ip"`
	Time_of_request uint64 `json:"time_of_request"`
}
