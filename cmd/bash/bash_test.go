package bash

import (
	"strings"
	"testing"
)

func TestBashBuild(t *testing.T) {
	bash := New().VolumeDir("/my/dir").Commands("echo banana", "echo batata").Build()

	if bash.container == nil {
		t.Error("should have a command")
	}

	if bash.container.VolumeDir != "/my/dir" {
		t.Error("should have the right volume dir")
	}

	if len(bash.commands) != len([]string{"echo banana", "echo batata"}) {
		t.Error("should have right command arguments")
	}

	expectedCommands := []string{"echo banana", "echo batata"}
	if strings.Compare(strings.Join(bash.commands, ""), strings.Join(expectedCommands, "")) != 0 {
		t.Error("should have right command arguments")
	}

}

func TestContainerArguments(t *testing.T) {
	workDir := "/tmp/build/"
	bash := New().VolumeDir("/my/dir").WorkDir(workDir).Commands("echo banana", "echo batata").Build()
	cmd := bash.ContainerArguments()

	expected := []string{
		"run", "--rm",
		"-v", "/my/dir:/tmp/build/",
		"-w", "/tmp/build/",
		"golang:1.11",
		"/bin/bash", "-c", "echo banana; echo batata; exit $?",
	}

	if len(cmd) != len(expected) {
		t.Error("should have the same length")
	}

	if strings.Compare(strings.Join(cmd, ""), strings.Join(expected, "")) != 0 {
		t.Error("should be equal")
	}

}
