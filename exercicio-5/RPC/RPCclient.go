package main

import (
	"fmt"
	"net/rpc"
	"time"
)

func main() {
	var reply string
	sampleSize := 10000

	// conecta ao servidor
	client, err := rpc.Dial("tcp", "localhost:1313")

	if err != nil {
		print(err)
	}

	defer client.Close()

	name := "Sofia"

	for i := 0; i < sampleSize; i++ {

		t1 := time.Now()
		err = client.Call("NinjaNameTransformer.TransformName", name, &reply)

		fmt.Println(time.Now().Sub(t1).Microseconds())

		if err != nil {
			print(err)
		}

	}

	// fmt.Println(reply)

}
