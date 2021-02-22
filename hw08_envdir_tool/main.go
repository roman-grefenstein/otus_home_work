package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("wrong count args")
	}

	envDirPath := os.Args[1]
	cmd := os.Args[2:]

	env, err := ReadDir(envDirPath)

	if err != nil {
		log.Fatalf("reading env dir error: %s", err)
	}

	os.Exit(RunCmd(cmd, env))
}
