package main

import (
	"fmt"
	"fswatch-command/command"
	"fswatch-command/config"
	"github.com/andreaskoch/go-fswatch"
)

func main() {
	configuration := config.ReadConfig()
	fmt.Println(configuration)

	checkIntervalInSeconds := 2
	for _, options := range configuration {
		fmt.Println(options)
		fileWatcher := fswatch.NewFileWatcher(options.File, checkIntervalInSeconds)
		fileWatcher.Start()

		for fileWatcher.IsRunning() {
			select {
			case <-fileWatcher.Modified():
				go func() {
					if options.On == "file-modified" {
						command.Cmd(options.Command, true)
					}
				}()
			case <-fileWatcher.Moved():
				go func() {
					if options.On == "file-moved" {
						command.Cmd(options.Command, true)
					}
				}()
			}
		}
	}
}
