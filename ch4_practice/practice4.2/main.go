package main

//!+
import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
	"strings"
)

// -sha384 任意文字列   : 任意文字列をSHA384でハッシュ化した結果
// -sha512 任意文字列   : 任意文字列をSHA512でハッシュ化した結果

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()

		if strings.Contains(str, "-sha") {
			buf := strings.SplitN(str, " ", 2)
			if buf[0] == "-sha384" {
				fmt.Printf("SHA384 = %x\n", sha512.Sum384([]byte(buf[1])))
			} else if buf[0] == "-sha512" {
				fmt.Printf("SHA512 = %x\n", sha512.Sum512([]byte(buf[1])))
			}
		} else {
			fmt.Printf("SHA256 = %x\n", sha256.Sum256([]byte(str)))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
