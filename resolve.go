/* This program resolves hostname to ip-address */

/* Licensed under GNU GPL V3 as available on http://www.gnu.org/licenses/gpl-3.0.html */

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: resolve hostname")
		os.Exit(1)
	}
	name := os.Args[1]
	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Resolution error")
		os.Exit(1)
	}
	fmt.Println("Resolved Address is :", addr.String())
	return
}
