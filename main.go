package main

import (
	"fmt"

	"github.com/mhv2408/bencoding/decode"
)

func main() {
	fmt.Println("Hello BIT-TORRENT")
	//fmt.Println(decode.DecodeToString("Added"))
	//res1 := decode.Decode("6:codingi120el6:Coding10:Challengese")
	// ll6:coding11:challengesel4:user4:testee
	//fmt.Println(res1)

	res2 := decode.Decode("ll6:coding10:challengesel4:user4:testee")
	fmt.Println(res2)

}
