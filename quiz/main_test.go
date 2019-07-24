package main_test

import (
	"os/exec"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMain(t *testing.T) {
	expected := `Problem #1: 1+1 =
Problem #2: 2+3 =
Problem #3: 13+17 =
You scored 0 out of 3.
`

	cmd := exec.Command("go", "run", ".")
	output, err := cmd.CombinedOutput()
	actual := string(output)

	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Fatal(diff)
	}
}
