package bash

import (
	"log"
	"strings"

	"github.com/joyciapp/joyci-core/cmd"
	"github.com/joyciapp/joyci-core/docker"
)

const defaultImage = "golang:1.11"
const defaultExecutable = "/bin/bash"

// Bash struct
type Bash struct {
	container *docker.Docker
	commands  []string
}

// New starts Bash builder
func New() Bash {
	b := Bash{}
	b.container = &docker.Docker{Image: defaultImage, Executable: defaultExecutable}
	return b
}

// Image sets a image name to the Object
func (b Bash) Image(imageName string) Bash {
	b.container.Image = imageName
	return b
}

// Executable sets an executable to run inside the container
func (b Bash) Executable(executablePath string) Bash {
	b.container.Executable = executablePath
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

// VolumeAndWorkDir sets both attributes the same value
func (b Bash) VolumeAndWorkDir(dir string) Bash {
	b.container.VolumeDir = dir
	b.container.WorkDir = dir
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

// Execute execute a set of commandsr
func (b Bash) Execute(commands ...string) (interface{}, error) {
	return b.Commands(commands...).Run()
}

// ContainerArguments to run inside docker
func (b Bash) ContainerArguments() []string {
	commandsToExec := strings.Join(append(b.commands, "exit $?"), "; ")
	container := b.container
	workDir := container.WorkDir

	return []string{
		"run", "--rm",
		"-v", container.VolumeDir + ":" + workDir,
		"-w", workDir,
		container.Image,
		container.Executable, "-c", commandsToExec,
	}
}

// Run Bash Commmands
func (b Bash) Run() (interface{}, error) {
	arguments := b.ContainerArguments()
	log.Println("cmd docker ", arguments)
	result, err := cmd.ExecCommand("docker", arguments...)
	return result, err
}
