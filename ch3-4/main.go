package main

import (
	"archive/zip"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	// 3.4 zipファイルをwebサーバーからダウンロード
	http.HandleFunc("/zip", zipHandler())
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func zipHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/zip")
		w.Header().Set("Content-Disposition", "attachment; filename=ascii_sample.zip")

		zipWriter := zip.NewWriter(w) // zip書き込み先をResponseWriterにする
		defer zipWriter.Close()

		fMap := map[string]string{
			"hoge.csv": "これは,CSV,です\n",
			"fuga.txt": "これはTXTです",
		}

		for k, v := range fMap {
			zw, err := zipWriter.Create(k)
			if err != nil {
				log.Fatal(err)
			}

			sr := strings.NewReader(v)
			io.Copy(zw, sr)
		}
	}
}
