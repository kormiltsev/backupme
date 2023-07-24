package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kormiltsev/backupme/compressers"
	"github.com/kormiltsev/backupme/encoders"
)

func main() {
	// encr or decr
	encrypt, decrypt := new(bool), new(bool)
	fileaddress := &flag.Args()[0]
	ext := strings.ToLower(filepath.Ext(*fileaddress))
	if ext != ".encrypted" {
		*encrypt = true
	} else {
		*decrypt = true
	}

	// compress file or folder
	var buf bytes.Buffer
	if err := compressers.Compresso(*fileaddress, &buf); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// encrypto
	massa, err := encoders.Encripto(buf.Bytes(), []byte(password))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// decrypto
	archive, err := encoders.Decripto(massa, []byte(password))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// decompress it
	buf2 := bytes.NewReader(archive)
	if err := compressers.Decompresso(buf2, "./result"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// notification
	fmt.Println("File decrypted at ./result/\n")

}
