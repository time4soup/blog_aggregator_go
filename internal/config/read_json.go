package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const configFileName = "/.gatorconfig.json"

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Read() Config {
	configFilePath := getConfigFilePath()

	configData, err := os.ReadFile(configFilePath)
	checkErr(err)

	config := &Config{}
	err = json.Unmarshal(configData, config)
	checkErr(err)

	return *config
}
func getConfigFilePath() string {
	homeDir, err := os.UserHomeDir()
	checkErr(err)
	return fmt.Sprint(homeDir, configFileName)
}
