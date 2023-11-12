package router

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/aixoio/amionline/server/config"
	"github.com/aixoio/amionline/server/data"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

func register_quit_request_handler(db_connecter *sql.DB, r *mux.Router, config_data *config.ConfigData, redis_client *redis.Client) {
	r.HandleFunc("/quit", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		ip := r.RemoteAddr
		key := fmt.Sprintf("quit:trys:%s", ip)

		ctx := context.Background()

		value_rstr := redis_client.Get(ctx, key)
		
		value_str := value_rstr.Val()

		value, err := strconv.Atoi(value_str)
		if err != nil {
			value = 0
		}
		value++

		ctx = context.Background()

		redis_client.SetEx(ctx, key, value, time.Hour)

		if value > 3 {
			w.WriteHeader(http.StatusForbidden)
			res := data.QuitData_Responce{
				Success: false,
				Error_msg: "Over 3 trys",
			}
			jsonres, err := json.Marshal(res)
			if err != nil {
				return
			}
			w.Write(jsonres)
			return
		}

	})
}
