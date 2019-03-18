package main

import (
	"os"
	"time"

	"github.com/joyciapp/joyci-core/cmd/bash"
	"github.com/joyciapp/joyci-core/cmd/git"
	"github.com/joyciapp/joyci-core/grpc/api"
)

var (
	pwd, _    = os.Getwd()
	appName   = "joyci-core"
	workDir   = "/tmp/build/"
	volumeDir = pwd + workDir
)

func runCommands() {
	git := git.New().VolumeDir(volumeDir).Build()
	git.Clone("git@github.com:joyciapp/joyci-core.git")

	bash := bash.New().VolumeDir(volumeDir + "/" + appName).WorkDir(workDir + "/" + appName).Build()
	bash.Execute(
		"echo running the tests suite inside the container",
		"go test ./... -v",
	)

	if _, err := os.Stat(volumeDir + "/" + appName); os.IsNotExist(err) {
		log.Println("should clone a git repository")
	}

}

func main() {
	// start the server
	go api.Serve()

	// execute client requests
	api.GitClone("git@github.com:joyciapp/joyci-core.git")

	api.ExecuteCommands(
		"echo running the tests suite inside the container",
		"go test ./...",
	)

	// waits until the commands finish
	time.Sleep(30 * time.Second)
	// After
	os.RemoveAll(volumeDir)
	//
	log.Println("program exit")
}
