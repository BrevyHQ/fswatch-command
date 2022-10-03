package watcher

import (
	"fswatch-command/config"
	"github.com/andreaskoch/go-fswatch"
)

type Instance struct {
	Watcher *fswatch.FileWatcher
	Config  config.WatcherConfig
}
