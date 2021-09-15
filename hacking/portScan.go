package hacking

import (
	"fmt"
	"net"
	"time"
)

func PortScan(host string, fromPort int, toPort int) []int {
	var findings []int
	for port := fromPort; port <= toPort; port++ {
		func(port int) {
			conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), time.Second)
			if err != nil {
				return
			}
			conn.Close()
			findings = append(findings, port)
		}(port)
	}
	return findings
}
