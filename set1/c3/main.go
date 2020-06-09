package main

import (
	"fmt"
	"strings"
)

func main() {
	keys := "abcdefghijklmnopqrstuvwxyz !\"#$%^'()*+,-./01223456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_{|}~"
}

func mostCommonLetter(str []byte) (char string) {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i := 0; i < len(str), i++ {
		for j := 0; j < len(alphabet), j++ {
			
		}
	}
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
