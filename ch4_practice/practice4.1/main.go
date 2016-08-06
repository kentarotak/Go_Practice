package main

//!+
import "crypto/sha256"
import "fmt"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	var test1 byte

	test1 = 0xFF

	a := uint16(test1)

	fmt.Printf("比較結果=%d\n", a)

	test := [2]byte{0xFF, 0xFF}

	fmt.Printf("比較結果=%d\n", compareSHA(c1, c2))
}

func compareSHA(c1 [32]byte, c2 [32]byte) int {
	var counter int
	for i := 0; i < len(c1); i++ {
		comp := byte(c1[i]) ^ byte(c2[i])
		counter += int(pc[comp])
	}
	return counter
}
