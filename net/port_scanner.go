package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

var website = flag.String("website", "scanme.nmap.org", "url to scann ports")

func main() {
	flag.Parse()
	var waitGroup sync.WaitGroup
	const maxPorts int = 65535

	for i := 0; i < maxPorts; i++ {
		defer waitGroup.Done()
		waitGroup.Add(1)
		go func(port int) {
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *website, port))
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("Port %d is open\n", port)
		}(i)
	}
	waitGroup.Wait()
}
