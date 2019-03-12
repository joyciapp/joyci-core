package git

import (
	"log"
	"os"
	"strings"
	"testing"
)

func TestGitBuild(t *testing.T) {
	commands := []string{"clone", "git@github.com:my-org/super-repo.git"}
	git := New().VolumeDir("/my/dir").Commands(commands...).Build()

	if git.container == nil {
		t.Error("should have a container")
	}

	if git.container.VolumeDir != "/my/dir" {
		t.Error("should have the right volume dir")
	}

	if len(git.commands) != len(commands) {
		t.Error("should have right command arguments")
	}

	if strings.Compare(strings.Join(git.commands, ""), strings.Join(commands, "")) != 0 {
		t.Error("should have right command arguments")
	}

}

func TestContainerArguments(t *testing.T) {
	commands := []string{"clone", "git@github.com:my-org/super-repo.git"}
	git := New().VolumeDir("/my/dir").Commands(commands...).Build()
	cmd := git.ContainerArguments()

	expected := []string{
		"run", "--rm",
		"-v", os.Getenv("HOME") + ":" + "/root", // to share ssh keys
		"-v", "/my/dir:/git",
		"alpine/git:1.0.7",
		"clone", "git@github.com:my-org/super-repo.git",
	}

	log.Println("cmd", cmd, len(cmd))
	log.Println("exp", expected, len(expected))

	if len(cmd) != len(expected) {
		t.Error("should have the same length")
	}

	if strings.Compare(strings.Join(cmd, ""), strings.Join(expected, "")) != 0 {
		t.Error("should be equal")
	}

}
