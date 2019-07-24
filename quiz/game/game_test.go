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
	expected, err := ioutil.ReadFile(filepath.Join("..", "testdata", t.Name()+".golden"))
	if err != nil {
		t.Fatal(err)
	}
	problems := []problems.Problem{
		{Question: "1+1", Answer: "2"},
		{Question: "is this a valid question?", Answer: "yes"},
		{Question: "one more?", Answer: "sure"},
	}
	answers := strings.NewReader("2\nyes\nnope")
	output := strings.Builder{}

	game.Play(&problems, answers, &output)

	if diff := cmp.Diff(string(expected), output.String()); diff != "" {
		t.Fatal(diff)
	}
}
