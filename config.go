package main

import (
	"code.google.com/p/gcfg"
	"errors"
	"os"
	"os/user"
)

const (
	KEY_ENV_NAME = "FILEPICKER_APIKEY"
	CONFIG_FILE  = ".fp"
)

func loadConfig() (Config, error) {
	var config Config

	usr, _ := user.Current()
	dir := usr.HomeDir

	apikey := os.Getenv(KEY_ENV_NAME)

	if apikey == "" {
		if err := gcfg.ReadFileInto(&config, dir+"/"+CONFIG_FILE); err != nil {
			return Config{}, err
		}
	} else {
		config.Filepicker.ApiKey = apikey
	}

	if config.Filepicker.ApiKey == "" {
		return Config{}, errors.New("ApiKey not found")
	}

	return config, nil

}
