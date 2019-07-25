package game

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/phss/gophercises/quiz/problems"
)

// Play a quiz using the set of problems.
func Play(problems *[]problems.Problem, timeLimit time.Duration, input io.Reader, output io.Writer) {
	score := 0

	c := make(chan int, 1)
	go play(c, problems, input, output, &score)

	select {
	case <-c:
	case <-time.After(timeLimit):
		fmt.Fprintln(output)
	}

	fmt.Fprintf(output, "You scored %d out of %d.\n", score, len(*problems))
}

func play(c chan int, problems *[]problems.Problem, input io.Reader, output io.Writer, score *int) {
	playerInputReader := bufio.NewReader(input)

	for i, problem := range *problems {
		fmt.Fprintf(output, "Problem #%d: %s = ", i+1, sanitise(problem.Question))
		answer, _, _ := playerInputReader.ReadLine()
		if sanitise(string(answer)) == sanitise(problem.Answer) {
			*score++
		}
	}
	c <- *score
}

func sanitise(s string) string {
	return strings.TrimSpace(s)
}
