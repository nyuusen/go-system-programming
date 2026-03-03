package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 2-1 ファイルに対するフォーマット出力
	f, err := os.Create("out.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(f, "数字:%d, 小数:%f, 文字列:%s", 123, 4.56, "abc")
}
