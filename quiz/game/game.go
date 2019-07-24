package game

import (
	"bufio"
	"fmt"
	"io"

	"github.com/phss/gophercises/quiz/problems"
)

// Play a quiz using the set of problems.
func Play(problems *[]problems.Problem, input io.Reader, output io.Writer) {
	userInput := bufio.NewReader(input)
	score := 0
	for i, problem := range *problems {
		fmt.Fprintf(output, "Problem #%d: %s =\n", i+1, problem.Question)
		userAnswer, _ := userInput.ReadString('\n')
		if userAnswer == problem.Answer+"\n" {
			score++
		}
	}

	fmt.Fprintf(output, "You scored %d out of %d.\n", score, len(*problems))
}
