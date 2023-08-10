package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	var reply string

	// conecta ao servidor
	client, err := rpc.Dial("tcp", "localhost:1313")

	if err != nil {
		print(err)
	}

	defer client.Close()

	name := "Sofia"

	err = client.Call("NinjaNameTransformer.TransformName", name, &reply)

	if err != nil {
		print(err)
	}

	fmt.Println(reply)

}
