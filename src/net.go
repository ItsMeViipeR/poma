package main

import (
	"fmt"

	"github.com/cakturk/go-netstat/netstat"
)

type PortInfo struct {
	IP    string
	Port  uint16
	Proto string
}

func GetOpenedPorts() ([]PortInfo, error) {
	var results []PortInfo
	seen := make(map[string]bool)

	tcpSocks, err := netstat.TCPSocks(func(s *netstat.SockTabEntry) bool {
		return s.State == netstat.Listen
	})

	if err != nil {
		return nil, err
	}

	udpSocks, err := netstat.UDPSocks(netstat.NoopFilter)

	if err != nil {
		return nil, err
	}

	for _, sock := range tcpSocks {
		key := fmt.Sprintf("tcp-%s:%d", sock.LocalAddr.IP, sock.LocalAddr.Port)

		if !seen[key] {
			seen[key] = true
			results = append(results, PortInfo{
				IP:    sock.LocalAddr.IP.String(),
				Port:  sock.LocalAddr.Port,
				Proto: "tcp",
			})
		}
	}

	for _, sock := range udpSocks {
		key := fmt.Sprintf("udp-%s:%d", sock.LocalAddr.IP, sock.LocalAddr.Port)

		if !seen[key] {
			seen[key] = true
			results = append(results, PortInfo{
				IP:    sock.LocalAddr.IP.String(),
				Port:  sock.LocalAddr.Port,
				Proto: "udp",
			})
		}
	}

	return results, nil
}
