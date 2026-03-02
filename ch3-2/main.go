package main

import (
	"crypto/rand"
	"io"
	"log"
	"os"
)

func main() {
	// 3.2 テスト用の適当なサイズのファイルを作成
	f, err := os.Create("bin")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	io.CopyN(f, rand.Reader, 1024)
}
