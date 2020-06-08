package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func decode(input []byte) ([]byte, error) {
	dst := make([]byte, hex.DecodedLen(len(input)))
	_, err := hex.Decode(dst, input)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func base64Encode(input []byte) []byte {
	eb := make([]byte, base64.StdEncoding.EncodedLen(len(input)))
	base64.StdEncoding.Encode(eb, input)

	return eb
}

func main() {
	message := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")

	result, err := decode(message)
	if err != nil {
		fmt.Printf("%s", err)
	}

	encodedMessage := base64Encode(result)

	fmt.Printf("%s", encodedMessage)
}
