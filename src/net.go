package main

import "github.com/cakturk/go-netstat/netstat"

func GetOpenedPorts() ([]uint16, []string, error) {
	socks, err := netstat.UDPSocks(netstat.NoopFilter)

	if err != nil {
		return nil, nil, err
	}

	var ports []uint16
	var addrs []string

	for _, sock := range socks {
		sock_local_addr := sock.LocalAddr

		ports = append(ports, sock_local_addr.Port)
		addrs = append(addrs, sock_local_addr.IP.String())
	}

	return ports, addrs, nil
}
