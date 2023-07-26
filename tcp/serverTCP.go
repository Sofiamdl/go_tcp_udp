package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var (
	dict map[string]string
)

func HandleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}(conn)
	for {
		req, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = conn.Write([]byte(TransformName(req) + "\n"))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func ServerTCP() {
	r, err := net.ResolveTCPAddr("tcp", "localhost:1313")
	if err != nil {
		fmt.Println(err)
	}

	ln, err := net.ListenTCP("tcp", r)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Servidor aguardando conexão...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go HandleConnection(conn)
	}
}

func TransformName(name string) string {
	newName := ""
	for i := 0; i < len(name); i++ {

		letter := strings.ToLower(string(name[i]))
		if i == 3 {
			newName += " "
		}
		newName += dict[letter]
	}
	return newName
}

func makeDic() map[string]string {
	myDict := make(map[string]string)
	myDict["a"] = "ka"
	myDict["b"] = "zu"
	myDict["c"] = "mi"
	myDict["d"] = "te"
	myDict["e"] = "ku"
	myDict["f"] = "lu"
	myDict["g"] = "ji"
	myDict["h"] = "ri"
	myDict["i"] = "ki"
	myDict["j"] = "zu"
	myDict["k"] = "me"
	myDict["l"] = "ta"
	myDict["m"] = "rin"
	myDict["n"] = "to"
	myDict["o"] = "mo"
	myDict["p"] = "no"
	myDict["q"] = "ke"
	myDict["r"] = "shi"
	myDict["s"] = "ari"
	myDict["t"] = "chi"
	myDict["u"] = "do"
	myDict["v"] = "ru"
	myDict["w"] = "mei"
	myDict["x"] = "na"
	myDict["y"] = "fu"
	myDict["z"] = "zi"
	return myDict
}

func main() {
	dict = makeDic()
	ServerTCP()
}
