package game_test

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/phss/gophercises/quiz/game"
	"github.com/phss/gophercises/quiz/problems"
)

func TestPlay(t *testing.T) {
	testTable := []struct {
		name     string
		problems []problems.Problem
		answers  []string
	}{
		{
			name: "normal game",
			problems: []problems.Problem{
				{Question: "1+1", Answer: "2"},
				{Question: "is this a valid question?", Answer: "yes"},
				{Question: "one more?", Answer: "sure"},
			},
			answers: []string{"2", "yes", "nope"},
		},
		{
			name: "sanitisation",
			problems: []problems.Problem{
				{Question: " 1+1           ", Answer: " 	2			"},
			},
			answers: []string{" 2  "},
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			expected, err := ioutil.ReadFile(filepath.Join("..", "testdata", t.Name()+".golden"))
			if err != nil {
				t.Fatal(err)
			}
			answers := strings.NewReader(strings.Join(testCase.answers, "\n"))
			output := strings.Builder{}

			game.Play(&testCase.problems, answers, &output)

			if diff := cmp.Diff(string(expected), output.String()); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
