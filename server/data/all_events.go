package data

type AllEvents_Responce struct {
	Cached bool `json:"cached"`
	Events []EventData `json:"events"`
	CachedAt uint64 `json:"cached_at"`
	Success bool `json:"success"`
	Error_msg string `json:"error_msg"`
}

type AllEvents_Cache struct {
	Events []EventData `json:"events"`
	SavedAt uint64 `json:"saved_at"`
}
