/* This program scans for all open ports from 1 to 65535 for the given host */

/* Licensed under GNU GPL V3 as available on http://www.gnu.org/licenses/gpl-3.0.html */

package main

import (
	"fmt"
	"net"
	"os"
)



func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: portscan ip-address")
		os.Exit(1)
	}

	host_str := os.Args[1]
	addr := net.ParseIP(host_str)
	if addr == nil {
		fmt.Println("Invalid ip-address")
		fmt.Println("Usage: portscan ip-address")
		os.Exit(1)
	}
	host := addr.String()

	fmt.Println("Port scanning started for " + host + ".")
	for port := 1; port <= 65535; port += 1 {
		connstr := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.Dial("tcp", connstr)
		if err == nil {
			fmt.Printf("TCP Port open : %d\n", port)
			conn.Close()
		}
	}
	fmt.Println("Port scanning finished.")
}

