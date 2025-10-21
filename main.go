package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mhv2408/bencoding/decode"
)

func main() {
	fmt.Println("Hello BIT-TORRENT")
	data, err := os.ReadFile("/Users/harshavardhanmirthinti/Downloads/world_will_idea_1301_librivox_archive.torrent")
	if err != nil {
		log.Fatal("Unable to open the file")
	}
	//fmt.Println(string(data))
	res := decode.Decode(string(data))

	if torrentData, ok := res.(map[string]any); ok {
		announceUrl := torrentData["announce"]
		infoMap := torrentData["info"]

		fmt.Println(announceUrl)
		if infoMapData, ok := infoMap.(map[string]any); ok {
			for Key := range infoMapData {
				fmt.Println(Key)
			}
		}

		//fmt.Println(infoMap)
	}

}
