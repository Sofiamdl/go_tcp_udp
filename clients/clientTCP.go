package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

var (
	names = []string{"Sofia", "Alex", "Felipo", "Elaine", "Aline", "Matheus"}
)

func ClientTCP(client string, sampleSize int) {

	r, err := net.ResolveTCPAddr("tcp", "localhost:1313")
	if err != nil {
		fmt.Println(err)
	}

	conn, err := net.DialTCP("tcp", nil, r)
	if err != nil {
		fmt.Println(err)
	}

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}(conn)

	for i := 0; i < sampleSize; i++ {

		req := client

		_, err = fmt.Fprintf(conn, req+"\n")
		if err != nil {
			fmt.Println(err)
		}
		t1 := time.Now()
		rep, err := bufio.NewReader(conn).ReadString('\n')

		fmt.Println(time.Now().Sub(t1).Microseconds())
		if err != nil {
			fmt.Println(err)
		}

		fmt.Print(req, " ", rep)
	}

}

func main() {
	n := 6
	sampleSize := 10
	for i := 0; i < n; i++ {
		ClientTCP(names[i], sampleSize)
	}
}
