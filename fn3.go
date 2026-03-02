package main

import (
	"io"
	"os"
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
	// h := func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application/zip")
	// 	w.Header().Set("Content-Disposition", "attachment; filename=ascii_sample.zip")

	// 	zipWriter := zip.NewWriter(w) // zip書き込み先をResponseWriterにする
	// 	defer zipWriter.Close()

	// 	fMap := map[string]string{
	// 		"hoge.csv": "これは,CSV,です\n",
	// 		"fuga.txt": "これはTXTです",
	// 	}

	// 	for k, v := range fMap {
	// 		zw, err := zipWriter.Create(k)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}

	// 		sr := strings.NewReader(v)
	// 		io.Copy(zw, sr)
	// 	}
	// }

	// http.HandleFunc("/zip", h)
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Fatal(err)
	// }

	// 3.5 CopyN
	// sr := strings.NewReader("ABCDEFGHIJK")
	// lReader := io.LimitReader(sr, 6)
	// io.Copy(os.Stdout, lReader)

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
	// programming.Seek(5, io.SeekStart)
	// io.Copy(os.Stdout, io.LimitReader(programming, 1)) // A
	stream = io.MultiReader(
		func() io.Reader {
			programming.Seek(5, io.SeekStart)
			defer programming.Seek(0, io.SeekStart)
			return io.LimitReader(programming, 1)
		}(),
		io.LimitReader(system, 1),
		io.LimitReader(computer, 1),
		func() io.Reader {
			programming.Seek(8, io.SeekStart)
			defer programming.Seek(0, io.SeekStart)
			return io.LimitReader(programming, 1)
		}(),
		func() io.Reader {
			programming.Seek(8, io.SeekStart)
			defer programming.Seek(0, io.SeekStart)
			return io.LimitReader(programming, 1)
		}(),
	)

	// ASCIIという文字列を出力する
	io.Copy(os.Stdout, stream)

}
