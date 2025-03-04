package main

import (
	"fmt"
	"github.com/cipher/des"
)

func main() {
	code, err := des.Code("ya sosal menya ebali", "key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Шифр: ", code)
	decode, err := des.Decode(code, "key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Дешифр: ", decode)
}
