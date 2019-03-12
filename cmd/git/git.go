package git

import (
	"log"
	"os"

	"github.com/joyciapp/joyci-core/cmd"
	"github.com/joyciapp/joyci-core/docker"
)

// Git struct
type Git struct {
	container *docker.Docker
	commands  []string
}

// New git
func New() Git {
	g := Git{}
	g.container = &docker.Docker{Image: "alpine/git:1.0.7"}
	return g
}

// VolumeDir to run inside the container
func (g Git) VolumeDir(volumeDir string) Git {
	g.container.VolumeDir = volumeDir
	return g
}

// Commands set commands to run inside the container
func (g Git) Commands(commands ...string) Git {
	g.commands = commands
	return g
}

// Build returns a new git
func (g Git) Build() *Git {
	return &g
}

// Clone appends commands to clone a git repository
func (g Git) Clone(repository string) {
	g.commands = []string{"clone", repository}
	g.Run()
}

// Run git command inside a container
func (g Git) Run() {
	arguments := g.ContainerArguments()
	log.Println("cmd docker ", arguments)
	cmd.ExecCommand("docker", arguments...)
}

// ContainerArguments ahsahjsh
func (g Git) ContainerArguments() []string {
	container := g.container
	arguments := []string{
		"run", "--rm",
		"-v", os.Getenv("HOME") + ":" + "/root", // to share ssh keys
		"-v", container.VolumeDir + ":" + "/git", //"-v", "$(pwd)/build/joyciapp:/git",
		container.Image,
	}

	for _, command := range g.commands {
		arguments = append(arguments, command)
	}

	return arguments
}
