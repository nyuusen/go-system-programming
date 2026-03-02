package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// 第3章 io.Reader
	// 3.1 ファイルのコピー
	// old file
	oldF, err := os.Open("old.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer oldF.Close()

	// new file
	newF, err := os.Create("new.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newF.Close()

	// copy
	io.Copy(newF, oldF)
}
