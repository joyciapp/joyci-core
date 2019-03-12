package cmd

import (
	"io"
	"log"
	"os"
	"os/exec"
)

// ExecCommand returns the command
func ExecCommand(command string, args ...string) (result interface{}, err error) {
	var cmd *exec.Cmd
	os.Setenv("PATH", "/bin:/usr/bin:/sbin:/usr/local/bin") // workaround
	cmd = exec.Command(command, args...)
	stdoutPipe, _ := cmd.StdoutPipe()
	stderrPipe, _ := cmd.StderrPipe()

	err = cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed %s\n%s\n", command, err)
		return nil, err
	}

	var errStdcapture error

	go func() {
		_, errStdcapture = io.Copy(os.Stdout, stdoutPipe)
	}()

	go func() {
		_, errStdcapture = io.Copy(os.Stderr, stderrPipe)
	}()

	err = cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Wait() failed with:\n%s\n", err)
		return nil, err
	}

	if errStdcapture != nil {
		log.Fatal("error on capture Stderr or Stdout")
		return nil, errStdcapture
	}

	return cmd, nil
}
