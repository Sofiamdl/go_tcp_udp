package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var (
	names = []string{"Sofia", "Alex", "Felipo", "Elaine", "Aline", "Matheus"}
)

func ClientTCP(i string) {

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

	//var mensagem string
	//fmt.Println("Digite sua mensagem por favor")

	req := i

	_, err = fmt.Fprintf(conn, req+"\n")
	if err != nil {
		fmt.Println(err)
	}

	rep, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(req, " ", rep)

}

func main() {
	n := 6
	for i := 0; i < n; i++ {
		ClientTCP(names[i])
	}
}
