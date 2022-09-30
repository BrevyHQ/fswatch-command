package config

import (
  "gopkg.in/yaml.v3"
  "io/ioutil"
  "os"
)

type Watcher struct {
  File     string
  On       string
  Command  string
  WithBash bool
}

func ReadConfig() map[string]Watcher {
  var configPath = os.Getenv("CONFIG_FILE_PATH")
  configFile, fileError := ioutil.ReadFile(configPath)

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
