package main

import (
	"flag"
	"fmt"
)

const version = "0.0.1"

func main() {
	flag.Usage = func() {
		fmt.Printf("joyci %s:\n", version)
		fmt.Println("--help print this screen")
		flag.PrintDefaults()
	}

	flag.Parse()
}
