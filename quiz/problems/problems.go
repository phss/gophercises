package problems

import (
	"encoding/csv"
	"fmt"
	"io"
)

// Problem represents a problem and a valid answer.
type Problem struct {
	Question string
	Answer   string
}

// Load problems from CSV formatted Reader.
func Load(reader io.Reader) ([]Problem, error) {
	lines, err := csv.NewReader(reader).ReadAll()
	if err != nil {
		return nil, err
	} else if len(lines) > 0 && len(lines[0]) != 2 {
		return nil, fmt.Errorf("incorrect number of columns in index %v", 0)
	}

	problems := make([]Problem, len(lines))
	for i, line := range lines {
		problems[i] = Problem{
			Question: line[0],
			Answer:   line[1],
		}
	}
	return problems, nil
}
