package main

import (
	"os"
	"os/exec"

	"github.com/godrei/logexample/logger"
)

func main() {
	cmd := exec.Command("go", "run", "./...")
	cmd.Dir = "/Users/godrei/Develop/go/src/github.com/godrei/logexample/command"

	logFilePth := "/Users/godrei/Develop/go/src/github.com/godrei/logexample/log.json"
	logFile, err := os.Create(logFilePth)
	if err != nil {
		panic(err)
	}
	cmd.Env = append(os.Environ(), "LOG_FILE="+logFilePth, "PROJECT_PATH=my_project.workspace")

	if err := logger.ServeLog(logFile); err != nil {
		panic(err)
	}

	if err := cmd.Start(); err != nil {
		panic(err)
	}
	if err := cmd.Wait(); err != nil {
		panic(err)
	}
}
