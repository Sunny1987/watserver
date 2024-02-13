package server

import (
	"encoding/json"
	"os"
)

// Config holds config.json properties
type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

// LoadConfiguration reads Config.json file
func LoadConfiguration(fileName string) (Config, error) {
	var config Config
	configFile, err := os.Open(fileName)
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParse := json.NewDecoder(configFile)
	err = jsonParse.Decode(&config)
	return config, err
}
