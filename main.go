package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/fsnotify/fsnotify"
)

func main() {
	playingWithFsnotify()
}

func runTests() {
	cmd := exec.Command("go", "test", "-v")

	stdoutStderr, err := cmd.CombinedOutput()
	fmt.Println("Test has run")
	if err != nil {
		fmt.Printf("Error has occurred: %s", err.Error())
	}
	fmt.Printf("%s\n", stdoutStderr)
	fmt.Println("========================================")
}

func playingWithFsnotify() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	// defer watcher.Close() - with error handling
	defer func(watcher *fsnotify.Watcher) {
		err := watcher.Close()
		if err != nil {
			log.Println(err)
		}
	}(watcher)

	// Start listening for events (through watcher.Events chan)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				//fmt.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("modified file:", event.Name)
					runTests()
				}
			case err := <-watcher.Errors:
				fmt.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(".")
	if err != nil {
		log.Fatal(err)
	}

	// Block main goroutine forever (until exit)
	<-make(chan struct{})
}
