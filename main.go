package main

import (
  "fswatch-command/config"
  "fswatch-command/watcher"
  "github.com/andreaskoch/go-fswatch"
  "os"
  "os/signal"
  "syscall"
)

func main() {
  configuration := config.ReadConfig()

  checkIntervalInSeconds := 2
  var watchers []watcher.Instance
  for _, options := range configuration {
    fileWatcher := fswatch.NewFileWatcher(options.File, checkIntervalInSeconds)
    fileWatcher.Start()
    instance := watcher.Instance{
      Watcher:  fileWatcher,
      Command:  options.Command,
      On:       options.On,
      WithBash: options.WithBash,
    }

    watchers = append(watchers, instance)
  }

  watcher.WatchForChanges(watchers)
}

func init() {
  c := make(chan os.Signal)
  signal.Notify(c, os.Interrupt, syscall.SIGTERM)
  go func() {
    <-c
    // Run Cleanup
    os.Exit(1)
  }()
}
