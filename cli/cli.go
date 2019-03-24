package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "0.0.1"

func parseAndRunSubCommand(command *flag.FlagSet, fn func()) {
	if os.Args[1] == command.Name() {
		command.Parse(os.Args[2:])
		if command.Parsed() {
			fn()
		}
	}
}

func versionSubcommand() {
	command := flag.NewFlagSet("version", flag.ExitOnError)
	parseAndRunSubCommand(command, func() {
		fmt.Println(version)
		os.Exit(0)
	})
}

func runSubcommand() {
	command := flag.NewFlagSet("run", flag.ExitOnError)
	pipelineFlag := command.String("pipeline", "", "Pipeline to run")

	parseAndRunSubCommand(command, func() {
		if *pipelineFlag == "" {
			fmt.Println("usage: joyci run --pipeline <pipeline-filename>")
			fmt.Println("examples:")
			fmt.Println("joyci run -pipeline=pipeline.go")
			fmt.Println("joyci run --pipeline=pipeline.go")
			os.Exit(1)
			return
		}

		fmt.Println("pipeline provided", *pipelineFlag)
		os.Exit(0)
	})
}

func main() {

	flag.Usage = func() {
		fmt.Println("usage: joyci <command> [<args>]")
		fmt.Println(" help print this screen")
		fmt.Println(" version print the current joyci version")
	}

	flag.Parse()

	if len(os.Args) == 1 || os.Args[1] == "help" {
		flag.Usage()
		os.Exit(0)
	}

	versionSubcommand()
	runSubcommand()
}
