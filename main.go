package main

import (
	"fmt"

	"github.com/mhv2408/bencoding/decode"
)

func main() {
	fmt.Println("Hello BIT-TORRENT")
	//fmt.Println(decode.DecodeToString("Added"))
	res := decode.Decode("6:codingi120e")
	fmt.Println(res)
}
