package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/phss/gophercises/quiz/problems"
)

func main() {
	defaultProblemsPath := flag.String("csv", "sample/problems.csv", "CSV with problem set")
	flag.Parse()

	problemsFile, _ := os.Open(*defaultProblemsPath)
	problems, _ := problems.Load(problemsFile)

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s =\n", i+1, problem.Question)
	}
	fmt.Printf("You scored 0 out of %d.\n", len(problems))
}
