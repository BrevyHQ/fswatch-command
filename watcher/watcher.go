package watcher

import (
	"fmt"
	"fswatch-command/command"
)

func WatchForChanges(watchers []Instance) {
	for anyWatchersRunning(watchers) {
		for _, instance := range watchers {
			select {
			case <-instance.Watcher.Modified():
				go func() {
					if instance.Config.On == "file-modified" {
						fmt.Println("File modified. Executing script.")
						command.Cmd(instance.Config.Command, instance.Config.Shell)
					}
				}()
			case <-instance.Watcher.Moved():
				go func() {
					if instance.Config.On == "file-moved" {
						fmt.Println("File moved. Executing script.")
						command.Cmd(instance.Config.Command, instance.Config.Shell)
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
