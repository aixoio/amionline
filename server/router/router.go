package router

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func Start(db_connecter *sql.DB) {
	r := mux.NewRouter()
	
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})

	http.ListenAndServe(":9090", r)
}
