package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type Watcher struct {
	File    string
	On      string
	Command string
}

func ReadConfig() map[string]Watcher {
	var configDir = os.Getenv("CONFIG_DIR")
	configFile, fileError := ioutil.ReadFile(configDir + "/config.yaml")

	if fileError != nil {
		panic(fileError)
	}

	data := make(map[string]Watcher)
	unmarshalError := yaml.Unmarshal(configFile, &data)

	if unmarshalError != nil {
		panic(unmarshalError)
	}

	return data
}
