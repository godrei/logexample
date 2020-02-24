package bitrise

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

// Step ...
type Step struct {
	Toolkit string
	CmdDir  string
	Inputs  []string
}

// Runner ...
type Runner struct {
	writer io.Writer
}

// New ...
func New(writer io.Writer) Runner {
	return Runner{
		writer: writer,
	}
}

// Run ...
func (r Runner) Run(steps []Step) error {
	envs := os.Environ()
	for i, step := range steps {
		r.writer.Write([]byte(fmt.Sprintf("Running step #%d\n", i+1)))

		if err := run(step, append(envs, step.Inputs...), r.writer); err != nil {
			return err
		}

		r.writer.Write([]byte("\n"))
	}
	return nil
}

func run(step Step, envs []string, logger io.Writer) error {
	cmd := exec.Command("go", "run", "main.go")
	cmd.Dir = step.CmdDir
	cmd.Env = envs
	cmd.Stderr, cmd.Stdout = logger, logger

	if err := cmd.Start(); err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}
