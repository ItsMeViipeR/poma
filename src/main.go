package main

import (
	"fmt"
	"os"
)

func showPorts() {
	ports, err := GetOpenedPorts()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	for _, p := range ports {
		fmt.Printf("[%s] %s:%d\n", p.Proto, p.IP, p.Port)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		showPorts()
	}
}
