package main

import (
	"fmt"
	"net"
	"os"
	"flag"
)

const (
	FRAME_LEN = 18
)

func main() {
	var host = flag.String("host", "localhost", "host")
	var port = flag.String("port", "3333", "port")
	flag.Parse()

	addr := *host+":"+*port
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on " + addr)

	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

func parseFrame(buf [] byte) int {
	mac := buf[0:6]
	fmt.Printf("MAC: %x:%x:%x:%x:%x:%x  RSSI %d", mac[0], mac[1], mac[2], mac[3], mac[4], mac[5], buf[7])
	return 0
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	var head int
	var tail int

	head = 0
	tail = 0
	buf := make([]byte, 1000)
	for{
		reqLen, err := conn.Read(buf[tail:])
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			conn.Close()
			return
		}
		tail += reqLen
		if tail < FRAME_LEN {
			continue
		}
		for head = 0; head < tail; head += 1 {
			if buf[head] != 0x54 {
				continue
			}
			if head + 1 < tail {
				if buf[head + 1] != 0x58 {
					continue
				} else {
					break
				}
			}
		}
		for ; head + FRAME_LEN < tail; head += FRAME_LEN {
			ret := parseFrame(buf)
			if ret  == -1 {
				conn.Close()
				return
			}
		}

		newBuf := make([]byte, 1000)
		copy(newBuf, buf[head:tail])
		buf = newBuf
		tail = 0
	}
}
