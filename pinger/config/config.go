package config

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/aixoio/amionline/logger"
)

type ConfigData struct {
	TargetIP string `json:"targetip"`
	IntervalSeconds uint64 `json:"intervalseconds"`
	ServerIP string `json:"serverip"`
}

func LoadConfig(path string) (*ConfigData, error) {
	logger.Info().Println("Loading config file", path)
	bytes, err := os.ReadFile(path)
	if err != nil {
		logger.Warn().Printf("Cannot load file %s trying env config\n", path)
		return LoadConfigFromEnv()
	}

	var config ConfigData
	json_err := json.Unmarshal(bytes, &config)
	if json_err != nil {
		return nil, json_err
	}

	return &config, nil
}

func LoadConfigFromEnv() (*ConfigData, error) {
	obj := ConfigData{}
	stu := reflect.ValueOf(&obj).Elem()
	typeof := stu.Type()
	for i := 0; i < stu.NumField(); i++ {
		env_name := fmt.Sprintf("ENV_%s", strings.ToUpper(typeof.Field(i).Name))
		env_val := os.Getenv(env_name)
		if env_val != "" {
			return nil, fmt.Errorf("empty env: %s", env_name)
		}
		field := stu.Field(i)

		switch field.Kind() {
		case reflect.String:
			field.SetString(env_val)
		case reflect.Bool:
			val, err := strconv.ParseBool(env_val)
			if err != nil {
				return nil, fmt.Errorf("failed to parse boolean value for %s: %v", env_name, err)
			}
			field.SetBool(val)
		default:
			return nil, fmt.Errorf("unsupported field type for %s: %v", env_name, field.Kind())
		}

	}

	return &obj, nil
}
