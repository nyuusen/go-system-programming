package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	// 3.5 CopyN
	sr := strings.NewReader("ABCDEFGHIJK")
	lReader := io.LimitReader(sr, 6)
	io.Copy(os.Stdout, lReader)
}
