package data

type QuitData struct {
	Pwd string `json:"pwd"`
}

type QuitData_Responce struct {
	Success bool `json:"success"`
	Error_msg string `json:"error_msg"`
}
