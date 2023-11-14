package events

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/aixoio/amionline/logger"
	"github.com/aixoio/amionline/pinger/config"
	"github.com/aixoio/amionline/pinger/data"
)

func Start(con *config.ConfigData) {
	logger.Info().Println("Starting events...")

	for {

		logger.Info().Printf("Pinging %s...\n", con.TargetIP)

		starttime := time.Now()

		resp, err := http.Get(con.TargetIP)
		if err != nil {
			logger.Warn().Println("Found intended error:", err.Error())
		} else {
			resp.Body.Close()
		}

		endtime := time.Now()

		success := true
		if err != nil {
			success = false
		}

		time_ms := endtime.UnixMilli() - starttime.UnixMilli()
		target_ip := con.TargetIP
		time_of_request := time.Now().UnixMilli()

		result := data.EventData{
			Success: success,
			Time_ms: float64(time_ms),
			Target_ip: target_ip,
			Time_of_request: uint64(time_of_request),
		}

		resbytes, err := json.Marshal(result)
		if err != nil {
			logger.Error().Println("Found error:", err.Error())
			continue
		}

		_, err = http.Post(con.ServerIP + "/api/log/event", "application/json", bytes.NewBuffer(resbytes))
		if err != nil {
			logger.Error().Println("Found error:", err.Error())
			continue
		}

		logger.Info().Printf("Pinged %s\n", con.TargetIP)

		time.Sleep(time.Duration(con.IntervalSeconds) * time.Second)

	}

}
