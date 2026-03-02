package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	// 3.6 ストリーム総集編
	// 制約
	// 1.ioパッケージと基本文法のみを使用する
	// 2.文字列リテラルは使用不可
	// 3.コメント部分以外は変更してはいけない
	var (
		computer    = strings.NewReader("COMPUTER")
		system      = strings.NewReader("SYSTEM")
		programming = strings.NewReader("PROGRAMMING")
	)

	var stream io.Reader
	a := io.NewSectionReader(programming, 5, 1)
	s := io.NewSectionReader(system, 0, 1)
	c := io.NewSectionReader(computer, 0, 1)
	i := io.NewSectionReader(programming, 8, 1)
	i2 := io.NewSectionReader(programming, 8, 1)
	stream = io.MultiReader(a, s, c, i, i2)

	// ASCIIという文字列を出力する
	io.Copy(os.Stdout, stream)
}
