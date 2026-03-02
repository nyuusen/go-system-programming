package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// 3.3 zipファイル書き込み
	zipF, err := os.Create("3-3.zip")
	if err != nil {
		log.Fatal(err)
	}
	zipWriter := zip.NewWriter(zipF)
	defer zipWriter.Close()

	zw, err := zipWriter.Create("text.txt")
	if err != nil {
		log.Fatal(err)
	}

	strReader := strings.NewReader("hoge")
	io.Copy(zw, strReader)
}
