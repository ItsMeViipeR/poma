package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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
		return
	}

	ports_info, err := GetOpenedPorts()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	activePorts := make(map[uint16]string)

	for _, p := range ports_info {
		ipMatch := p.IP == "localhost" || strings.HasPrefix(p.IP, "192.168.1.") || p.IP == "127.0.0.1" || p.IP == "0.0.0.0"

		if ipMatch {
			activePorts[p.Port] = p.IP
		}
	}

	for _, arg := range args {
		port, err := strconv.ParseUint(arg, 10, 16)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		targetPort := uint16(port)

		ip, found := activePorts[targetPort]

		if !found {
			fmt.Printf("Port %d not in use\n", targetPort)
		}

		pid, err := GetPIDByPort(ip, targetPort)

		if err != nil {
			fmt.Fprintf(os.Stderr, "PID error: %v\n", err)
			continue
		}

		err = KillProcessByPID(int(pid))

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while closing: %v\n", err)
			continue
		}
	}
}
