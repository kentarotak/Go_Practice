package archivelib

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func Decompress(file, dest string) error {
	var err error
	if strings.Contains(file, "zip") {
		err = unzip(file, dest)
	}

	return err
}

func unzip(file, dest string) error {
	r, err := zip.OpenReader(file)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.CopyN(os.Stdout, rc, 68)
		if err != nil {
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println()
	}
	return err
}
