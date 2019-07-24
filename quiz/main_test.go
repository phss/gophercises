package main_test

import (
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMain(t *testing.T) {
	testTable := []struct {
		name string
		args []string
	}{
		{name: "default params", args: []string{}},
		{name: "short example", args: []string{"-csv", "sample/short.csv"}},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			expected, err := ioutil.ReadFile(filepath.Join("testdata", t.Name()+".golden"))
			if err != nil {
				t.Fatal(err)
			}

			args := []string{"run", "."}
			args = append(args, testCase.args...)
			cmd := exec.Command("go", args...)
			actual, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(string(expected), string(actual)); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
