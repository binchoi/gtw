package watcher

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/fsnotify/fsnotify"
)

func runTests(path string) {
	cmd := exec.Command("go", "test", "-v", path)

	stdoutStderr, err := cmd.CombinedOutput()
	fmt.Println("Test has run")
	defaultColor := Green
	if err != nil {
		fmt.Printf(Red+"Error has occurred: %s"+Red, err.Error()) // TODO: refactor
		defaultColor = Red
	}
	fmt.Printf(defaultColor+"%s\n"+defaultColor, stdoutStderr)
	fmt.Println("========================================")
}

func StartWatcher(path string) {
	log.Printf("> Test directory path: %s\n", path)

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
					runTests(path)
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
