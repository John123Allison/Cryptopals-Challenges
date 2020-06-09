package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	var x string = "1c0111001f010100061a024b53535009181c"
	y := "686974207468652062756c6c277320657965"

	f, err := hex.DecodeString(x)
	k, err := hex.DecodeString(y)

	fmt.Printf("%v\n", len(f))
	fmt.Printf("%v\n", len(y))

	if err != nil {
		fmt.Println("Error")
	}

	a := []byte(f)
	b := []byte(k)

	fmt.Printf("%v", hex.EncodeToString(xor(a, b)))
}

func xor(x []byte, y []byte) []byte {
	if len(x) != len(y) {
		fmt.Println("Error: buffers are inequal length")
		return nil
	}

	buf := make([]byte, len(x))

	for i, _ := range x {
		buf[i] = x[i] ^ y[i]
	}

	return buf
}
