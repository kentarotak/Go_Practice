package main

import (
	"fmt"
)

const (
	KB = 1000
	MB = KB * KB
	GB = MB * MB
	TB = GB * GB
	PB = TB * TB
	EB = PB * PB
	//ZB = EB * EB
	//YB = ZB * ZB
)

func main() {
	fmt.Printf("%d\n", 10<<3+10<<1)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%d\n", GB)
}
