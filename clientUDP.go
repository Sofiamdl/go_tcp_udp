package main

import (
	"fmt"
	"net"
	"os"
)

var (
	names = []string{"Sofia", "Alex", "Felipo", "Elaine", "Aline", "Matheus"}
)

func udpClient(i string) {
	rep := make([]byte, 1024)

	addr, err := net.ResolveUDPAddr("udp", "localhost:1313")
	if err != nil {
		fmt.Println(err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println(err)
	}

	defer func(conn net.UDPConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}(*conn)

	if err != nil {
		fmt.Println(err)
	}

	req := []byte(i)

	_, err = conn.Write(req)

	if err != nil {
		fmt.Println(err)
	}

	_, _, err = conn.ReadFromUDP(rep)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(i), " ", string(rep), "\n")

}

func main() {
	n := 6
	for i := 0; i < n; i++ {
		udpClient(names[i])
	}
}
