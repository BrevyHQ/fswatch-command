package watcher

import (
  "fswatch-command/command"
)

func WatchForChanges(watchers []Instance) {
  for anyWatchersRunning(watchers) {
    for _, instance := range watchers {
      select {
      case <-instance.Watcher.Modified():
        go func() {
          if instance.On == "file-modified" {
            command.Cmd(instance.Command, true)
          }
        }()
      case <-instance.Watcher.Moved():
        go func() {
          if instance.On == "file-moved" {
            command.Cmd(instance.Command, true)
          }
        }()
      }
    }
  }
}

func anyWatchersRunning(watchers []Instance) bool {
  anyRunning := true
  for _, instance := range watchers {
    if !instance.Watcher.IsRunning() {
      anyRunning = false
      break
    }
  }

  return anyRunning
}
