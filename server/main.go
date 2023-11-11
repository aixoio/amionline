package main

import (
	"github.com/aixoio/amionline/server/db"
	"github.com/aixoio/amionline/server/router"
)

func main() {
	d, _ := db.Connect("root:root@tcp(127.0.0.1:3306)/amionline")
	defer db.Disconnect(d)
	router.Start()
}
