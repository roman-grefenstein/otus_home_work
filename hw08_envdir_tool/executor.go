package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	envParams := make([]string, len(env))
	for envName, envValue := range env {
		if envValue.NeedRemove {
			os.Unsetenv(envName)
			continue
		}
		_, ok := os.LookupEnv(envName)
		if ok {
			os.Unsetenv(envName)
		}
		envParams = append(envParams, fmt.Sprintf("%s=%s", envName, envValue.Value))
	}

	command := exec.Command(cmd[0], cmd[1:]...) //nolint:gosec
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	command.Env = append(os.Environ(), envParams...)

	var exitError *exec.ExitError
	if err := command.Run(); err != nil {
		if errors.As(err, &exitError) {
			return exitError.ExitCode()
		}
	}
	return 0
}
