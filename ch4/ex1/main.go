package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(1&i)
	}
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	bc := diffBits(c1, c2)
	fmt.Println(bc)
}

func diffBits(c1, c2 [32]byte) int {
	res := 0
	for i := 0; i < 32; i++ {
		diff := int(pc[(c1[i] ^ c2[i])])
		res += diff
		fmt.Printf("%x,%x,%x\n", c1[i], c2[i], diff)
	}
	return res
}
