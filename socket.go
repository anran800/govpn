package main

import (
	"net"
)

const GOVPN_PORT = 1194

type linkSocket struct {
	sd *net.UDPConn
}

func isValidIpOrDns(addr string) bool {
	return net.ParseIP(addr) != nil
}

func isValidIpv4Port(port int) bool {
	return port > 0 && port < 65536
}
