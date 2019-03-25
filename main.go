package main

import (
	"io/ioutil"

	"github.com/joyciapp/joyci-core/runner"
)

func main() {
	script, _ := ioutil.ReadFile("./examples/pipeline_go")
	job := runner.Job{
		BuildPath: "/tmp/build",
		Script:    string(script),
	}

	runner.Run(job)
}
