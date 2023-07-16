// package main

// import (
// 	"fmt"
// 	"net"
// 	"strings"
// )

// var (
// 	dict map[string]string
// )

// func udpServer() {
// 	req := make([]byte, 1024)

// 	addr, err := net.ResolveUDPAddr("udp", "localhost:1313")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	conn, err := net.ListenUDP("udp", addr)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println("Servidor UDP aguardando requests...")

// 	for {
// 		_, addr, err := conn.ReadFromUDP(req)
// 		if err != nil {
// 			fmt.Println(err)
// 		}

// 		_, err = conn.WriteTo([]byte(TransformName(string(req))), addr)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 	}

// }

// func TransformName(name string) string {
// 	newName := ""
// 	for i := 0; i < len(name); i++ {

// 		letter := strings.ToLower(string(name[i]))
// 		if i == 3 {
// 			newName += " "
// 		}
// 		newName += dict[letter]
// 	}
// 	return newName
// }

// func makeDic() map[string]string {
// 	myDict := make(map[string]string)
// 	myDict["a"] = "ka"
// 	myDict["b"] = "zu"
// 	myDict["c"] = "mi"
// 	myDict["d"] = "te"
// 	myDict["e"] = "ku"
// 	myDict["f"] = "lu"
// 	myDict["g"] = "ji"
// 	myDict["h"] = "ri"
// 	myDict["i"] = "ki"
// 	myDict["j"] = "zu"
// 	myDict["k"] = "me"
// 	myDict["l"] = "ta"
// 	myDict["m"] = "rin"
// 	myDict["n"] = "to"
// 	myDict["o"] = "mo"
// 	myDict["p"] = "no"
// 	myDict["q"] = "ke"
// 	myDict["r"] = "shi"
// 	myDict["s"] = "ari"
// 	myDict["t"] = "chi"
// 	myDict["u"] = "do"
// 	myDict["v"] = "ru"
// 	myDict["w"] = "mei"
// 	myDict["x"] = "na"
// 	myDict["y"] = "fu"
// 	myDict["z"] = "zi"
// 	return myDict
// }

// func main() {
// 	dict = makeDic()
// 	udpServer()
// }
