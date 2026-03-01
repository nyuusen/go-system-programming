package main

import (
	"archive/zip"
	"io"
	"log"
	"net/http"
	"strings"
)

// 第3章
func fn3() {
	// // 第3章 io.Reader
	// // 3.1 ファイルのコピー
	// // old file
	// oldF, err := os.Open("tmp/old.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer oldF.Close()

	// // new file
	// newF, err := os.Create("tmp/new.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer newF.Close()

	// // copy
	// io.Copy(newF, oldF)

	// // 3.2 テスト用の適当なサイズのファイルを作成
	// f, err := os.Create("tmp/bin")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	// io.CopyN(f, rand.Reader, 1024)

	// 3.3 zipファイル書き込み
	// zipF, err := os.Create("tmp/3-3.zip")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// zipWriter := zip.NewWriter(zipF)
	// defer zipWriter.Close()

	// zw, err := zipWriter.Create("text.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// strReader := strings.NewReader("hoge")
	// io.Copy(zw, strReader)

	// 3.4 zipファイルをwebサーバーからダウンロード
	h := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/zip")
		w.Header().Set("Content-Disposition", "attachment; filename=ascii_sample.zip")

		zipWriter := zip.NewWriter(w)
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

	http.HandleFunc("/zip", h)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
