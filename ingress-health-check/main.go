package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
