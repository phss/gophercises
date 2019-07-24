package main

import (
	"flag"
	"os"

	"github.com/phss/gophercises/quiz/game"

	"github.com/phss/gophercises/quiz/problems"
)

func main() {
	defaultProblemsPath := flag.String("csv", "sample/problems.csv", "CSV with problem set")
	flag.Parse()

	problemsFile, _ := os.Open(*defaultProblemsPath)
	problems, _ := problems.Load(problemsFile)

	game.Play(&problems, os.Stdin, os.Stdout)
}
