package config

import (
	"encoding/json"
	"os"

	"github.com/aixoio/amionline/logger"
)

type ConfigData struct {
	Mysqldb string `json:"mysqldb"`
	Redisurl string `json:"redisurl"`
	Quitpwd string `json:"quitpwd"`
	UncachedRoutes bool `json:"uncached_routes"`
	AutoClearCache bool `json:"auto_clear_cache"`
}

func LoadConfig(path string) (*ConfigData, error) {
	logger.Info().Println("Loading config file", path)
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config ConfigData
	json_err := json.Unmarshal(bytes, &config)
	if json_err != nil {
		return nil, json_err
	}

	return &config, nil
}
