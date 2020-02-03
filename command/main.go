package main

import (
	"os"
	"time"

	"github.com/godrei/logexample/log"
)

func main() {
	projectPth := os.Getenv("PROJECT_PATH")
	logPth := os.Getenv("LOG_FILE")

	f, err := os.Create(logPth)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	log.SetOut(f)

	log.Info("Building Xcode project")
	time.Sleep(1 * time.Second)
	log.Command("xcode build --project " + projectPth)
	time.Sleep(2 * time.Second)
	log.Info("Building Xcode project")
	time.Sleep(1 * time.Second)
	log.Command("xcode build --project " + projectPth)
}
