package main

import (
	"crypto/rand"
	"crypto/sha1"
	"log"
	"net/url"

	"github.com/mhv2408/bencoding/encode"
)

func EncodeInfoHash(input map[string]any) string {

	// encode the hash map to string
	encodedHashMap := encode.Encode(input)

	// 1. Encode the map to SHA-1
	hasher := sha1.New() // new encoder
	_, err := hasher.Write([]byte(encodedHashMap))
	if err != nil {
		log.Fatal("cannot write the hash: ", err)
	}
	sha := hasher.Sum(nil) // getting the SHA-1 hash as a byte slice

	// URL encode the sha1  string
	URLEncodedSHA1 := url.QueryEscape(string(sha))
	return URLEncodedSHA1
}
func EncodePeerID() string {
	ClientCode := "-HV1001-"
	length := 12 // 12 bytes of random
	random := make([]byte, length)
	rand.Read(random)
	Instance := url.QueryEscape(string(random))
	return ClientCode + Instance

}

func GetLength(files []any) int {
	Total := 0
	for _, f := range files {
		fileMap := f.(map[string]any)
		length := fileMap["length"].(int)
		Total += length
	}
	return Total
}
