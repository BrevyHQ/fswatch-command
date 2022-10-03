package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type WatcherConfig struct {
	File    string
	On      string
	Command string
	Shell   bool
}

func ReadConfig() map[string]WatcherConfig {
	var configPath = os.Getenv("CONFIG_FILE_PATH")
	configFile, fileError := ioutil.ReadFile(configPath)

	if fileError != nil {
		panic(fileError)
	}

	data := make(map[string]WatcherConfig)
	unmarshalError := yaml.Unmarshal(configFile, &data)

	if unmarshalError != nil {
		panic(unmarshalError)
	}

	return data
}
