package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(user string) error {
	cfg.CurrentUserName = user

	cfgData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	homePath, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	cfgFilePath := fmt.Sprint(homePath, configFileName)
	err = os.WriteFile(cfgFilePath, cfgData, 0666)
	if err != nil {
		return err
	}

	return nil
}
