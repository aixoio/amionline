package config

import (
	"encoding/json"
	"os"
)

type ConfigData struct {
	Db string `json:"db"`
}

func LoadConfig(path string) (*ConfigData, error) {
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
