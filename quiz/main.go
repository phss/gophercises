package main

import (
	"flag"
	"fmt"
	"os"

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

	problems, _ := problems.Load(problemsFile)

	game.Play(&problems, os.Stdin, os.Stdout)
}
