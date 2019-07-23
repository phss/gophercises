package questions

import (
	"encoding/csv"
	"io"
)

// Question represents a question and a valid answer.
type Question struct {
	Question string
	Answer   string
}

// Load questions from CSV formatted Reader.
func Load(reader io.Reader) ([]Question, error) {
	csv := csv.NewReader(reader)
	lines, err := csv.ReadAll()

	if err != nil {
		return nil, err
	}

	questions := make([]Question, len(lines))
	for i, line := range lines {
		questions[i] = Question{
			Question: line[0],
			Answer:   line[1],
		}
	}
	return questions, nil
}
