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

	torrentData, ok := res.(map[string]any)
	if !ok {
		log.Fatal("torrentData does not match the pattern map[string]any")
	}

	announceRaw, ok := torrentData["announce"]
	if !ok {
		log.Fatal("missing announce URL")
	}

	announceURL, ok := announceRaw.(string)
	if !ok {
		log.Fatal("annoincrURL is not a string")
	}

	infoRaw, ok := torrentData["info"]
	if !ok {
		log.Fatal("infoMap does not exist")
	}

	infoMap, ok := infoRaw.(map[string]any)
	if !ok {
		log.Fatal("infoMap is not a map")
	}
	infoHash := EncodeInfoHash(infoMap)

	peerId := EncodePeerID()

	filesRaw, ok := infoMap["files"]
	if !ok {
		log.Fatal("files not in infoMap")
	}
	files, ok := filesRaw.([]any)
	if !ok {
		log.Fatal("files is not a list")
	}
	left := GetLength(files)

	url := fmt.Sprintf("%s?info_hash=%s&peer_id=%s&port=6881&uploaded=0&downloaded=0&left=%d&compact=1", announceURL, infoHash, peerId, left)

	fmt.Printf("Announce URL: <%s>", url)

}
