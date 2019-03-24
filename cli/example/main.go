package main

import (
	cli "github.com/joyciapp/joyci-core/grpc/api"
)

func main() {
	cli.GitClone("git@github.com:joyciapp/joyci-core.git")
}
