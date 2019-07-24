package questions

import (
	"encoding/csv"
	"fmt"
	"io"
)

// Question represents a question and a valid answer.
type Question struct {
	Question string
	Answer   string
}

// Load questions from CSV formatted Reader.
func Load(reader io.Reader) ([]Question, error) {
	lines, err := csv.NewReader(reader).ReadAll()
	if err != nil {
		return nil, err
	} else if len(lines) > 0 && len(lines[0]) != 2 {
		return nil, fmt.Errorf("incorrect number of columns in index %v", 0)
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
