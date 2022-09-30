package watcher

import "github.com/andreaskoch/go-fswatch"

type Instance struct {
  Watcher  *fswatch.FileWatcher
  WithBash bool
  Command  string
  On       string
}
