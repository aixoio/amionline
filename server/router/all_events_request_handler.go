package router

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/aixoio/amionline/server/data"
	"github.com/aixoio/amionline/server/logger"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

func register_all_events_request_handler(r *mux.Router, db_connecter *sql.DB, redis_client *redis.Client) {
	r.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		logger.Info().Printf("Handling request to %s from %s\n", r.URL.Path, r.RemoteAddr)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Headers", "*")
        w.Header().Set("Access-Control-Allow-Origin", "*")

		ctx := context.Background()
		redis_v := redis_client.Get(ctx, "events:all")
		redis_s := redis_v.Val()

		var all_cached data.AllEvents_Cache
		err := json.Unmarshal([]byte(redis_s), &all_cached)
		canusecache := true
		if err != nil {
			canusecache = false
		}

		if canusecache {
			w.Header().Set("Redis-Cache", "HIT")
			result := data.AllEvents_Responce{
				Cached: true,
				Events: all_cached.Events,
				CachedAt: all_cached.SavedAt,
				Success: true,
				Error_msg: "",
			}
			jsonres, err := json.Marshal(result)
			if err != nil {
				logger.Error().Printf("Found error: %s\n", err.Error())
				return
			}
			w.Write(jsonres)
			return
		}

		w.Header().Set("Redis-Cache", "MISS")

		tx, err := db_connecter.Begin()
		if err != nil {
			logger.Error().Printf("Found error: %s\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			res := data.AllEvents_Responce{
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

		rows, err := tx.Query("SELECT id, success, time_ms, target_ip, time_of_request FROM events")
		if err != nil {
			logger.Error().Printf("Found error: %s\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			res := data.AllEvents_Responce{
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
				res := data.AllEvents_Responce{
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

		data_to_cache := data.AllEvents_Cache{
			Events: events,
			SavedAt: uint64(time.Now().UnixMilli()),
		}
		cache_bytes, err := json.Marshal(data_to_cache)
		if err != nil {
			logger.Error().Printf("Found error: %s\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			res := data.AllEvents_Responce{
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

		ctx = context.Background()
		redis_client.SetEx(ctx, "events:all", string(cache_bytes), time.Hour)

		w.WriteHeader(http.StatusOK)
		responce := data.AllEvents_Responce{
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
