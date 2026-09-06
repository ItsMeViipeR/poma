package main

import (
	"fmt"
	"os"

	"github.com/shirou/gopsutil/v3/net"
)

func KillProcessByPID(pid int) error {
	proc, err := os.FindProcess(pid)

	if err != nil {
		return fmt.Errorf("process not found %d: %w", pid, err)
	}

	err = proc.Kill()

	if err != nil {
		return fmt.Errorf("cannot kill process %d: %w", pid, err)
	}

	return nil
}

func GetPIDByPort(targetIP string, targetPort uint16) (int32, error) {
	connections, err := net.Connections("all")

	if err != nil {
		return 0, err
	}

	for _, conn := range connections {
		if conn.Laddr.Port == uint32(targetPort) {
			if conn.Laddr.IP == targetIP || targetIP == "0.0.0.0" || conn.Laddr.IP == "0.0.0.0" {
				return conn.Pid, nil
			}
		}
	}

	return 0, fmt.Errorf("no process found for %s:%d", targetIP, targetPort)
}
