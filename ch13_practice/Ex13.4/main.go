package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {

	cmd := exec.Command("bzip2", "-c")
	cmd.Stdin = os.Stdin

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

}
