package events

import (
	"sync"
	"time"

	"github.com/aixoio/amionline/logger"
	"github.com/aixoio/amionline/pinger/config"
)

func Start(con *config.ConfigData) {
	logger.Info().Println("Starting events...")
	
	wg := sync.WaitGroup{}

	for {

		go ping(con, &wg)

		time.Sleep(time.Duration(con.IntervalSeconds) * time.Second)

	}

	wg.Wait()

}
