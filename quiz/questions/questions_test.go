package questions_test

import (
	"strings"
	"testing"

	"github.com/phss/gophercises/quiz/questions"
)

func TestLoad(t *testing.T) {
	expectedQuestions := []questions.Question{
		{Question: "1+1", Answer: "2"},
		{Question: "is this a valid question?", Answer: "yes"},
		{Question: "one more?", Answer: "sure"},
	}
	actualQuestions := questions.Load(strings.NewReader(`
1+1,2
is this a valid question?,yes
one more?,sure`))

	if len(actualQuestions) != len(expectedQuestions) {
		t.Errorf("Mismatch number of errores. Expected %v but got %v.", len(expectedQuestions), len(actualQuestions))
	}
	for i := range actualQuestions {
		if actualQuestions[i] != expectedQuestions[i] {
			t.Errorf("On index %v expected %v but got %v", i, expectedQuestions[i], actualQuestions[i])
		}

	}
}
