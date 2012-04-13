/* This program scans for all open ports from 1 to 65535 for the given host */

/* Licensed under GNU GPL V3 as available on http://www.gnu.org/licenses/gpl-3.0.html */

package main

import (
	"fmt"
	"net"
	"os"
)

var ic = make(chan int, 1000)

func port_scan(host string, port int) {
	connstr := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("tcp", connstr)
	if err == nil {
		fmt.Printf("TCP Port open : %d\n", port)
		conn.Close()
	}
	ic <- 1
}

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

	port := 1
	counter := 1
	for port <= 65535 {
		/* running go routines max 1000 at a time */
		for counter = 1; counter < 1000; counter += 1 { // 65535 // max 1018
			go port_scan(host, port)
			port += 1
			if port > 65535 {
				break
			}
		}
		for chan_counter := 1; chan_counter < counter; chan_counter += 1 {
			<-ic
		}
	}
	fmt.Println("Port scanning finished.")
}

