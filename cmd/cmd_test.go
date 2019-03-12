package cmd

import (
	"testing"
)

func TestExecCommand(t *testing.T) {
	var result interface{}
	var err error

	result, err = ExecCommand("cat", "cmd.go")
	if result == nil || err != nil {
		t.Error("should handle one line command call")
	}

	result, err = ExecCommand("./script.sh")
	if result == nil || err != nil {
		t.Error("should execute a script")
	}
}
