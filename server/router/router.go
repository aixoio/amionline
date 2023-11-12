package router

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

func Start(db_connecter *sql.DB, redis_client *redis.Client) {
	r := mux.NewRouter()
	
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})

	register_log_event_request_handler(r, db_connecter)

	http.ListenAndServe(":9090", r)
}
