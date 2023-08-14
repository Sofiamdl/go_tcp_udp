package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

var (
	names = []string{"Sofia", "Alex", "Felipo", "Elaine", "Aline", "Matheus", "Miguel",
		"Davi",
		"Gabriel",
		"Arthur",
		"Lucas",
		"Matheus",
		"Pedro",
		"Guilherme",
		"Gustavo",
		"Rafael",
		"Felipe",
		"Bernardo",
		"Enzo",
		"Nicolas"}
)

func udpClient(client string, sampleSize int) {
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

	for i := 0; i < sampleSize; i++ {

		if err != nil {
			fmt.Println(err)
		}

		req := []byte(client)

		t1 := time.Now()
		_, err = conn.Write(req)

		_, _, err = conn.ReadFromUDP(rep)
		fmt.Println(time.Now().Sub(t1).Microseconds())

		if err != nil {
			fmt.Println(err)
		}

		//Print Name
		//fmt.Print(string(i), " ", string(rep), "\n")
	}

}

func main() {
	randNum := rand.Intn(20)
	sampleSize := 10000
	udpClient(names[randNum], sampleSize)

}
