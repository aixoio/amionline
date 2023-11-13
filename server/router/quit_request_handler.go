package router

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aixoio/amionline/server/config"
	"github.com/aixoio/amionline/server/data"
	"github.com/aixoio/amionline/server/db/mysql"
	"github.com/aixoio/amionline/server/logger"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

func register_quit_request_handler(db_connecter *sql.DB, r *mux.Router, config_data *config.ConfigData, redis_client *redis.Client) {
	r.HandleFunc("/quit", func(w http.ResponseWriter, r *http.Request) {
		logger.Info().Printf("Handling request to %s from %s\n", r.URL.Path, r.RemoteAddr)
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

		var input data.QuitData
		input_err := json.NewDecoder(r.Body).Decode(&input)
		if input_err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			res := data.QuitData_Responce{
				Success: false,
				Error_msg: "Cannot read request body",
			}
			jsonres, err := json.Marshal(res)
			if err != nil {
				return
			}
			w.Write(jsonres)
			return
		}

		if input.Pwd == config_data.Quitpwd {

			ctx := context.Background()
			redis_client.Del(ctx, key)

			w.WriteHeader(http.StatusOK)
			res := data.QuitData_Responce{
				Success: true,
				Error_msg: "",
			}
			jsonres, err := json.Marshal(res)
			if err != nil {
				return
			}
			w.Write(jsonres)

			mysql.Disconnect(db_connecter)

			os.Exit(0)
			return
		}

		w.WriteHeader(http.StatusUnauthorized)
		res := data.QuitData_Responce{
			Success: false,
			Error_msg: "Incorrect password",
		}
		jsonres, err := json.Marshal(res)
		if err != nil {
			return
		}
		w.Write(jsonres)

	})
}
