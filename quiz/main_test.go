package main_test

import (
	"os/exec"
	"testing"
)

func TestMain(t *testing.T) {
	expected := "Hello world\n"

	cmd := exec.Command("go", "run", ".")
	output, err := cmd.CombinedOutput()
	actual := string(output)

	if err != nil {
		t.Fatal(err)
	}
	if actual != expected {
		t.Fatalf("expected %s, got %s", expected, actual)
	}
}
