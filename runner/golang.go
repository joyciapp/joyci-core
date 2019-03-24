package runner

import (
	"github.com/joyciapp/joyci-core/cmd/bash"
)

const (
	defaultImage      = "golang:1.11"
	defaultExecutable = "/usr/local/go/bin/go"
	defaultFilename   = "pipeline.go"
)

// Job represents a job to run inside a Runner
type Job struct {
	id         string
	Image      string
	Executable string
	BuildPath  string
	Filename   string
	Script     string
}

// Run a given script
func Run(job Job) {
	if job.Image == "" {
		job.Image = defaultImage
	}

	if job.Executable == "" {
		job.Executable = defaultExecutable
	}

	if job.Filename == "" {
		job.Filename = defaultFilename
	}

	bash := bash.New()
	runner := bash.Image(job.Image).VolumeAndWorkDir(job.BuildPath).Build()
	runner.Execute(
		"echo '"+job.Script+"' > "+job.BuildPath+"/"+job.Filename,
		"go run "+job.Filename,
	)
}
