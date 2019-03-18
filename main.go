package main

import (
	"os"

	"github.com/joyciapp/joyci-core/cmd/bash"
	"github.com/joyciapp/joyci-core/cmd/git"
)

func main() {
	volumeDir, _ := os.Getwd()
	workDir := "/tmp/build/"

	git := git.New().VolumeDir(volumeDir + "/tmp/build/joyciapp").Build()
	git.Clone("git@github.com:joyciapp/joyci-core.git")

	bash := bash.New().VolumeDir(volumeDir).WorkDir(workDir).Commands(
		"echo running the tests suite inside the container",
		"go test ./... -v",
	)
	bash.Build().Run()
}
