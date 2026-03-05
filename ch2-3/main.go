package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	source := map[string]string{
		"Hello": "World",
	}

	// map to json
	b, err := json.Marshal(source)
	if err != nil {
		fmt.Printf("json marshal error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// multi writer
	mr := io.MultiWriter(w, os.Stdout)

	// gzip圧縮
	gw := gzip.NewWriter(mr)
	defer gw.Flush()
	if _, err = gw.Write(b); err != nil {
		fmt.Printf("gzip write error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
