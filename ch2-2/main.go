package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	// 2-2 CSV出力
	cw := csv.NewWriter(os.Stdout)
	defer cw.Flush()

	s := [][]string{
		{"abc", "def", "ghi"},
		{"jkl", "mno", "pqr"},
	}

	for _, v := range s {
		if err := cw.Write(v); err != nil {
			log.Fatal(err)
		}
	}
}
