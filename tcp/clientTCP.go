package main

import (
	"bufio"
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

		t1 := time.Now()
		_, err = fmt.Fprintf(conn, req+"\n")
		if err != nil {
			fmt.Println(err)
		}

		_, err := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(time.Now().Sub(t1).Microseconds())

		if err != nil {
			fmt.Println(err)
		}

		//Name print
		//fmt.Print(req, " ", rep)
	}

}

func main() {
	randNum := rand.Intn(20)
	sampleSize := 10000
	ClientTCP(names[randNum], sampleSize)
}
