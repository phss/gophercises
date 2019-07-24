package problems_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/phss/gophercises/quiz/problems"
)

func TestLoad(t *testing.T) {
	expected := []problems.Problem{
		{Question: "1+1", Answer: "2"},
		{Question: "is this a valid question?", Answer: "yes"},
		{Question: "one more?", Answer: "sure"},
	}
	actual, err := problems.Load(strings.NewReader(`
1+1,2
is this a valid question?,yes
one more?,sure`))

	if err != nil {
		t.Errorf("got error %v", err)
	}
	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Error(diff)
	}
}

func TestLoad_incorrectCsv(t *testing.T) {
	_, err := problems.Load(strings.NewReader(`
wrong format
missing a column`))

	if err.Error() != "incorrect number of columns in index 0" {
		t.Errorf("Expected failure but got %v", err)
	}
}

func TestLoad_errorReader(t *testing.T) {
	_, err := problems.Load(failingReader{"should fail"})

	if err.Error() != "should fail" {
		t.Errorf("Expected failure but got %v", err)
	}
}

type failingReader struct {
	failure string
}

func (r failingReader) Read(p []byte) (int, error) {
	return 0, errors.New(r.failure)
}
