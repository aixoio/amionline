package router

import (
	"database/sql"
	"net/http"

	"github.com/aixoio/amionline/server/config"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

func Start(db_connecter *sql.DB, redis_client *redis.Client, config_data *config.ConfigData) {
	r := mux.NewRouter()
	
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})

	register_log_event_request_handler(r, db_connecter)
	register_quit_request_handler(db_connecter, r, config_data, redis_client)
	register_last_20_events_request_handler(r, db_connecter, redis_client)
	register_all_events_request_handler(r, db_connecter, redis_client)

	http.ListenAndServe(":9090", r)
}
