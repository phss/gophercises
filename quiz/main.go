package main

import (
	"fmt"
	"os"

	"github.com/phss/gophercises/quiz/questions"
)

func main() {
	defaultProblemsPath := "sample/short.csv"
	problemsFile, _ := os.Open(defaultProblemsPath)

	problems, _ := questions.Load(problemsFile)

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s =\n", i+1, problem.Question)
	}
	fmt.Printf("You scored 0 out of %d.\n", len(problems))
}
