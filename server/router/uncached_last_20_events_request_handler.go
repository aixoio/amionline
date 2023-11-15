package router

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/aixoio/amionline/logger"
	"github.com/aixoio/amionline/server/data"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

func register_uncached_last_20_events_request_handler(r *mux.Router, db_connecter *sql.DB, redis_client *redis.Client) {
	r.HandleFunc("/api/events/uncached/last/20", func(w http.ResponseWriter, r *http.Request) {
		logger.Info().Printf("Handling request to %s from %s\n", r.URL.Path, r.RemoteAddr)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Headers", "*")
        w.Header().Set("Access-Control-Allow-Origin", "*")

		tx, err := db_connecter.Begin()
		if err != nil {
			logger.Error().Printf("Found error: %s\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			res := data.Last20Events_Responce{
				Success: false,
				Error_msg: "Cannot start database transaction",
			}
			jsonres, err := json.Marshal(res)
			if err != nil {
				logger.Error().Printf("Found error: %s\n", err.Error())
				return
			}
			w.Write(jsonres)
			return
		}
		defer tx.Commit()

		rows, err := tx.Query("SELECT id, success, time_ms, target_ip, time_of_request FROM events ORDER BY id DESC LIMIT 20")
		if err != nil {
			logger.Error().Printf("Found error: %s\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			res := data.Last20Events_Responce{
				Success: false,
				Error_msg: "Cannot load data from database",
			}
			jsonres, err := json.Marshal(res)
			if err != nil {
				logger.Error().Printf("Found error: %s\n", err.Error())
				return
			}
			w.Write(jsonres)
			return
		}
		defer rows.Close()

		events := []data.EventData{}

		for rows.Next() {

			var (
				id uint64
				success bool
				time_ms float64
				target_ip string
				time_of_request uint64
			)

			err := rows.Scan(&id, &success, &time_ms, &target_ip, &time_of_request)
			if err != nil {
				logger.Error().Printf("Found error: %s\n", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				res := data.Last20Events_Responce{
					Success: false,
					Error_msg: "Cannot load data from database",
				}
				jsonres, err := json.Marshal(res)
				if err != nil {
					logger.Error().Printf("Found error: %s\n", err.Error())
					return
				}
				w.Write(jsonres)
				return
			}

			events = append(events, data.EventData{
				Id: id,
				Success: success,
				Time_ms: time_ms,
				Target_ip: target_ip,
				Time_of_request: time_of_request,
			})

		}

		data_to_cache := data.Last20Events_Cache{
			Events: events,
			SavedAt: uint64(time.Now().UnixMilli()),
		}
		cache_bytes, err := json.Marshal(data_to_cache)
		if err != nil {
			logger.Error().Printf("Found error: %s\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			res := data.Last20Events_Responce{
				Success: false,
				Error_msg: "Cannot store to cache",
			}
			jsonres, err := json.Marshal(res)
			if err != nil {
				logger.Error().Printf("Found error: %s\n", err.Error())
				return
			}
			w.Write(jsonres)
			return
		}

		ctx := context.Background()
		redis_client.SetEx(ctx, "events:last:20", string(cache_bytes), time.Hour / 2)

		w.WriteHeader(http.StatusOK)
		responce := data.Last20Events_Responce{
			Cached: false,
			Events: events,
			CachedAt: 0,
			Success: true,
			Error_msg: "",
		}
		jsonres, err := json.Marshal(responce)
		if err != nil {
			logger.Error().Printf("Found error: %s\n", err.Error())
			return
		}
		w.Write(jsonres)


	})
}
