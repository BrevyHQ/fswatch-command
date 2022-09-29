package main

import (
	"fswatch-command/config"
	"fswatch-command/watcher"
	"github.com/andreaskoch/go-fswatch"
)

func main() {
  configuration := config.ReadConfig()

  checkIntervalInSeconds := 2
  var watchers []watcher.Instance
  for _, options := range configuration {
    fileWatcher := fswatch.NewFileWatcher(options.File, checkIntervalInSeconds)
    fileWatcher.Start()
    withConfig := watcher.Instance{Watcher: fileWatcher, Command: options.Command, On: options.On}
    watchers = append(watchers, withConfig)
  }

  watcher.WatchForChanges(watchers)
}
