package runner

import (
	"io/ioutil"
	"testing"
)

func TestRun(t *testing.T) {
	script, _ := ioutil.ReadFile("../examples/pipeline.go")

	job := Job{
		BuildPath: "/tmp/build",
		Script:    string(script),
	}

	Run(job)
}
