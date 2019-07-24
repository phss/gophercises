package main_test

import (
	"io/ioutil"
	"os/exec"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMain(t *testing.T) {
	cmd := exec.Command("go", "run", ".")
	actual, err := cmd.CombinedOutput()
	expected, _ := ioutil.ReadFile("testdata/TestMain.golden")

	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(string(expected), string(actual)); diff != "" {
		t.Fatal(diff)
	}
}
