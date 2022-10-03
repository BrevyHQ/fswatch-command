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
		instance := watcher.Instance{
			Watcher: fileWatcher,
			Config:  options,
		}

		watchers = append(watchers, instance)
	}

	watcher.WatchForChanges(watchers)
}
