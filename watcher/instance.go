package watcher

import "github.com/andreaskoch/go-fswatch"

type Instance struct {
	Watcher *fswatch.FileWatcher
	Command string
	On      string
}
