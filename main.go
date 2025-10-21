package main

import (
	"fmt"

	"github.com/mhv2408/bencoding/decode"
	"github.com/mhv2408/bencoding/encode"
)

func main() {
	fmt.Println("Hello BIT-TORRENT")
	//fmt.Println(decode.DecodeToString("Added"))
	//res1 := decode.Decode("6:codingi120el6:Coding10:Challengese")
	// ll6:coding11:challengesel4:user4:testee
	//fmt.Println(res1)

	//res2 := decode.Decode("ll6:coding10:challengesel4:user4:testeed6:codingl4:user4:testee")

	//res3 := decode.Decode("d6:codingl4:user4:testee")
	//res3 := "d6:codingl4:user4:testee"

	//fmt.Println(res2)

	res3 := decode.Decode("d17:Coding Challengesd6:Rating7:Awesome8:website:20:codingchallenges.fyiee")
	fmt.Println(res3)

	// fmt.Println(encode.Encode("harsha"))

	// fmt.Println(encode.Encode(123))
	// fmt.Println(encode.GetValue(res3))
	//fmt.Println(encode.Encode(res2))
	fmt.Println(encode.Encode(res3))
}
