package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	envs := make(Environment, len(files))

	for _, file := range files {
		if file.IsDir() || strings.Contains(file.Name(), "=") {
			continue
		}

		if file.Size() == 0 {
			envs[file.Name()] = EnvValue{NeedRemove: true}
			continue
		}

		f, err := os.Open(dir + "/" + file.Name())
		if err != nil {
			return nil, err
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		scanner.Scan()
		line := scanner.Text()
		line = strings.TrimRight(line, "\t ")
		line = strings.ReplaceAll(line, "\x00", "\n")
		envs[file.Name()] = EnvValue{Value: line, NeedRemove: len(line) == 0}
	}
	return envs, nil
}
