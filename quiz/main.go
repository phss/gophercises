package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/phss/gophercises/quiz/game"

	"github.com/phss/gophercises/quiz/problems"
)

func main() {
	csvFilename := flag.String("csv", "sample/problems.csv", "CSV with problem set")
	flag.Parse()

	problemsFile, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Printf("Failed to open file '%s'\n", *csvFilename)
		os.Exit(1)
	}

	problems, err := problems.Load(problemsFile)
	if err != nil {
		fmt.Printf("Invalid file format: %s\n", err)
		os.Exit(1)
	}

	game.Play(&problems, 30*time.Second, os.Stdin, os.Stdout)
}
