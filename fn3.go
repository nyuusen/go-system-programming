package main

import (
	"crypto/rand"
	"io"
	"log"
	"os"
)

// 第3章
func fn3() {
	// 第3章 io.Reader
	// 3.1 ファイルのコピー
	// old file
	oldF, err := os.Open("tmp/old.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer oldF.Close()

	// new file
	newF, err := os.Create("tmp/new.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newF.Close()

	// copy
	io.Copy(newF, oldF)

	// 3.2 テスト用の適当なサイズのファイルを作成
	f, err := os.Create("tmp/bin")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	io.CopyN(f, rand.Reader, 1024)

	// 3.3
}
