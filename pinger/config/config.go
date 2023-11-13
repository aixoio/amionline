package config

import (
	"encoding/json"
	"os"

	"github.com/aixoio/amionline/logger"
)

type ConfigData struct {
	TargetIP string `json:"target_ip"`
	IntervalSeconds uint64 `json:"interval_seconds"`
	ServerIP string `json:"server_ip"`
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
