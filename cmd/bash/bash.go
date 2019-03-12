package bash

import (
	"strings"

	"github.com/joyciapp/joyci-core/cmd"
	"github.com/joyciapp/joyci-core/docker"
)

// Bash struct
type Bash struct {
	container *docker.Docker
	commands  []string
}

// New re
func New() Bash {
	b := Bash{}
	b.container = &docker.Docker{Image: "golang:1.11", Executable: "/bin/bash"}
	return b
}

// VolumeDir to run inside the container
func (b Bash) VolumeDir(volumeDir string) Bash {
	b.container.VolumeDir = volumeDir
	return b
}

// WorkDir to run inside the container
func (b Bash) WorkDir(workDir string) Bash {
	b.container.WorkDir = workDir
	return b
}

// Commands to run inside the container
func (b Bash) Commands(commands ...string) Bash {
	b.commands = commands
	return b
}

// Build returns a new bash
func (b Bash) Build() *Bash {
	return &b
}

// ContainerArguments ahsahjsh
func (b Bash) ContainerArguments() []string {
	commandsToExec := strings.Join(append(b.commands, "exit $?"), "; ")
	container := b.container
	workDir := container.WorkDir + "/" + "$(git rev-parse --short=7 HEAD)"

	return []string{
		"run", "--rm",
		"-v", container.VolumeDir + ":" + workDir,
		"-w", workDir,
		container.Image,
		container.Executable, "-c", commandsToExec,
	}
}

// Run Bash Commmands
func (b Bash) Run() {
	cmd.ExecCommand("docker", b.ContainerArguments()...)
}
