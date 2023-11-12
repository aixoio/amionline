package router

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/aixoio/amionline/server/data"
	"github.com/gorilla/mux"
)

func register_log_event_request_handler(r *mux.Router, db_connecter *sql.DB) {
	r.HandleFunc("/log/event", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Headers", "*")
        w.Header().Set("Access-Control-Allow-Origin", "*")

		tx, err := db_connecter.Begin()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			res := data.EventData_Responce{
				Success: false,
				Error_msg: "Cannot open database transaction",
			}
			jsonres, err := json.Marshal(res)
			if err != nil {
				return
			}
			w.Write(jsonres)
		}
		defer tx.Commit()
		
		var event_data data.EventData

		json_err := json.NewDecoder(r.Body).Decode(&event_data)
		if json_err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			res := data.EventData_Responce{
				Success: false,
				Error_msg: "Cannot parse data",
			}
			jsonres, err := json.Marshal(res)
			if err != nil {
				return
			}
			w.Write(jsonres)
		}

		res, err := tx.Exec("INSERT INTO `events` (`success`, `time_ms`, `target_ip`, `time_of_request`) VALUES (?, ?, ?, ?)",
							event_data.Success, event_data.Time_ms, event_data.Target_ip, event_data.Time_of_request)
		rows_added, extra_err := res.RowsAffected()
		if err != nil || rows_added != 1 || extra_err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			res := data.EventData_Responce{
				Success: false,
				Error_msg: "Cannot store data in database",
			}
			jsonres, err := json.Marshal(res)
			if err != nil {
				return
			}
			w.Write(jsonres)
		}

		w.WriteHeader(http.StatusOK)
		responce := data.EventData_Responce{
			Success: true,
		}
		jsonres, err := json.Marshal(responce)
		if err != nil {
			return
		}
		w.Write(jsonres)

	})
}
