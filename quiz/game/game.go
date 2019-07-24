package game

import (
	"bufio"
	"fmt"
	"io"

	"github.com/phss/gophercises/quiz/problems"
)

// Play a quiz using the set of problems.
func Play(problems *[]problems.Problem, input io.Reader, output io.Writer) {
	playerInputReader := bufio.NewReader(input)
	score := 0
	for i, problem := range *problems {
		fmt.Fprintf(output, "Problem #%d: %s = ", i+1, problem.Question)
		answer, _, _ := playerInputReader.ReadLine()
		if string(answer) == problem.Answer {
			score++
		}
	}

	fmt.Fprintf(output, "You scored %d out of %d.\n", score, len(*problems))
}
