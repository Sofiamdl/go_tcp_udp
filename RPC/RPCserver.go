package main

import (
	"fmt"
	"net"
	"net/rpc"
	"strings"
)

type NinjaNameTransformer string

var (
	dict map[string]string
)

func (n *NinjaNameTransformer) TransformName(nameToChange string, reply *string) error {
	*reply = TransformName(nameToChange)
	return nil
}

func RPCServer() {
	// instancia do transformador do nome ninja
	ninjaNameTransformer := new(NinjaNameTransformer)

	//cria novo servidor rpc
	server := rpc.NewServer()

	//registra nome ninja
	err := server.RegisterName("NinjaNameTransformer", ninjaNameTransformer)
	if err != nil {
		print(err)
	}

	//cria listener tcp
	ln, err := net.Listen("tcp", "localhost:1313")
	if err != nil {
		print(err)
	}

	defer func(ln net.Listener) {
		var err = ln.Close()
		if err != nil {
			print(err)
		}
	}(ln)

	// aguarda por invocações
	fmt.Println("Aguardando conexões...")
	server.Accept(ln)
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
	RPCServer()

}
