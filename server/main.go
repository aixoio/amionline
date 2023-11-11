package main

import (
	"fmt"

	"github.com/aixoio/amionline/server/config"
	"github.com/aixoio/amionline/server/db"
	"github.com/aixoio/amionline/server/router"
)

func main() {
	c, _ := config.LoadConfig("env.json")
	fmt.Println(c.Msg)
	d, _ := db.Connect("root:root@tcp(127.0.0.1:3306)/amionline")
	defer db.Disconnect(d)
	router.Start()
}
