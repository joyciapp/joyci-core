package main

import (
	"log"
	"os"

	"github.com/joyciapp/joyci-core/cmd/bash"
	"github.com/joyciapp/joyci-core/cmd/git"
)

func main() {
	pwd, _ := os.Getwd()
	appName := "joyci-core"
	workDir := "/tmp/build/"
	volumeDir := pwd + workDir

	git := git.New().VolumeDir(volumeDir).Build()
	git.Clone("git@github.com:joyciapp/joyci-core.git")

	bash := bash.New().VolumeDir(volumeDir+"/"+appName).WorkDir(workDir+"/"+appName).Commands(
		"echo running the tests suite inside the container",
		"go test ./...",
	)
	bash.Build().Run()

	if _, err := os.Stat(volumeDir + "/" + appName); os.IsNotExist(err) {
		log.Println("should clone a git repository")
	}

	// After
	os.RemoveAll(volumeDir)

}
