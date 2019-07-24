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
	}{
		{name: "default params"},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			expected, err := ioutil.ReadFile(filepath.Join("testdata", t.Name()+".golden"))
			if err != nil {
				t.Fatal(err)
			}

			cmd := exec.Command("go", "run", ".")
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
