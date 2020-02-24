package main

import (
	"os"

	"github.com/godrei/logexample/bitrise"
)

const (
	exampleStep = "/Users/godrei/Develop/go/src/github.com/godrei/logexample/xcodebuildstep"
	logPth      = "/Users/godrei/Develop/go/src/github.com/godrei/logexample/bitrise-run.log"
)

func initLog(pth string) (*os.File, error) {
	log, err := os.Create(pth)
	if err != nil {
		return nil, err
	}
	return log, nil
}

func readSteps() []bitrise.Step {
	return []bitrise.Step{
		bitrise.Step{
			Toolkit: "go",
			CmdDir:  exampleStep,
			Inputs:  []string{"PROJECT_PATH=./tmp.xcodeproj"},
		},
		bitrise.Step{
			Toolkit: "go",
			CmdDir:  exampleStep,
			Inputs:  []string{"PROJECT_PATH=./tmp.xcworkspace"},
		},
	}
}

func main() {
	log, err := initLog(logPth)
	if err != nil {
		panic(err)
	}
	defer log.Close()

	steps := readSteps()

	if err := bitrise.New(log).Run(steps); err != nil {
		panic(err)
	}
}
